import { Component, OnInit, inject } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { firstValueFrom } from 'rxjs';

import { ApolloQueryResult } from '@apollo/client/core';
import { MutationResult, QueryRef } from 'apollo-angular';

import { CheckboxModule } from 'primeng/checkbox';

import {
  TaskStatus,
  HomePage_ListTasksGQL,
  HomePage_ListTasksQuery,
  HomePage_ListTasksQueryVariables,
  HomePage_CompleteTaskGQL,
  HomePage_CompleteTaskMutation,
} from '../../../gen/graphql-codegen/schema';

import { AddTaskButtonComponent } from './add-task-button/add-task-button.component';

import { Task } from '../../interfaces/pages/home-page';

import { ErrorService } from '../../services/error.service';
import { NotificationService } from '../../services/notification.service';

import * as utils from '../../utils';

import { environment } from '../../../environments/environment';

@Component({
  selector: 'app-home-page',
  standalone: true,
  imports: [CommonModule, FormsModule, CheckboxModule, AddTaskButtonComponent],
  templateUrl: './home-page.component.html',
  styleUrl: './home-page.component.css',
})
export class HomePageComponent implements OnInit {
  private listTasksGQL = inject(HomePage_ListTasksGQL);
  private listTasksQuery: QueryRef<
    HomePage_ListTasksQuery,
    HomePage_ListTasksQueryVariables
  >;

  private completeTaskGQL = inject(HomePage_CompleteTaskGQL);

  private errorService = inject(ErrorService);
  private notificationService = inject(NotificationService);

  public tasks: Task[] = [];
  public checkedTaskIDs: string[] = [];
  public isTasksReady: boolean = false;
  public isTaskCompleting: boolean = false;

  constructor() {
    this.listTasksQuery = this.listTasksGQL.watch({
      status: TaskStatus.Uncompleted,
      first: environment.gql.defaultEdgeCountPerPage,
    });
  }

  public ngOnInit(): void {
    this.initTasks();
  }

  private async initTasks(refetch: boolean = false): Promise<void> {
    {
      const extractTasks = (_query: HomePage_ListTasksQuery): Task[] => {
        return (
          _query.tasks.edges
            // CompleteTask を実行した際、完了した Task 単体のキャッシュは更新される（status は TaskStatus.Completed になる）が、
            // 該当 Task がクエリ結果のキャッシュから除外されるわけではないことを考慮する必要がある。
            .filter((_edge) => _edge.node.status === TaskStatus.Uncompleted)
            .map((_edge) => _edge.node)
        );
      };

      let result: ApolloQueryResult<HomePage_ListTasksQuery>;

      try {
        if (refetch) {
          result = await this.listTasksQuery.refetch();
        } else {
          result = await this.listTasksQuery.result();
        }
      } catch (e) {
        this.errorService.handle(e);
        return;
      }

      this.tasks = extractTasks(result.data);
      this.isTasksReady = true;

      while (result.data.tasks.pageInfo.hasNextPage) {
        try {
          result = await this.listTasksQuery.fetchMore({
            variables: {
              after: result.data.tasks.pageInfo.endCursor,
            },
          });
        } catch (e) {
          this.errorService.handle(e);
          return;
        }

        this.tasks.push(...extractTasks(result.data));
      }
    }
  }

  public async onChangeTaskCheckbox(task: Task): Promise<void> {
    if (!this.checkedTaskIDs.includes(task.id)) {
      return;
    }

    this.isTaskCompleting = true;

    {
      let result: MutationResult<HomePage_CompleteTaskMutation>;
      {
        try {
          result = await firstValueFrom(
            this.completeTaskGQL.mutate({
              input: {
                id: task.id,
              },
            }),
          );
        } catch (e) {
          this.errorService.handle(e);
          this.isTaskCompleting = false;
          return;
        }
      }

      const err = result.data!.completeTask.error;
      if (err !== undefined && err !== null) {
        switch (err.__typename) {
          case 'BadRequestError':
            this.notificationService.badRequest(err.message);
            break;
          default:
            this.errorService.handle(new Error(err.message));
        }
        this.isTaskCompleting = false;
        return;
      }
    }

    await utils.sleep(500);

    this.tasks = this.tasks.filter((_task) => _task.id !== task.id);
    this.isTaskCompleting = false;
  }

  public onCompleteAddTask(): void {
    this.initTasks(true);
  }
}
