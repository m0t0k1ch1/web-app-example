import { Component, OnInit, inject } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { firstValueFrom } from 'rxjs';

import { ApolloQueryResult } from '@apollo/client/core';
import { QueryRef } from 'apollo-angular';

import { CheckboxModule } from 'primeng/checkbox';
import { RippleModule } from 'primeng/ripple';

import {
  CompleteTaskGQL,
  ListTasksQuery,
  ListTasksQueryVariables,
  ListTasksGQL,
  Task,
  TaskStatus,
} from '../../../gen/graphql-codegen/schema';

import { NotificationService } from '../../services/notification.service';

import * as utils from '../../utils';

@Component({
  selector: 'app-home-page',
  standalone: true,
  imports: [CommonModule, FormsModule, CheckboxModule, RippleModule],
  templateUrl: './home-page.component.html',
  styleUrl: './home-page.component.css',
})
export class HomePageComponent implements OnInit {
  private listTasksGQL = inject(ListTasksGQL);
  private completeTaskGQL = inject(CompleteTaskGQL);

  private notificationService = inject(NotificationService);

  private listTasksQuery: QueryRef<ListTasksQuery, ListTasksQueryVariables>;

  public tasks: Task[] = [];
  public checkedTaskIDs: string[] = [];
  public isTasksReady: boolean = false;
  public isTaskCompleting: boolean = false;

  public isAddTaskButtonHovered: boolean = false;

  constructor() {
    this.listTasksQuery = this.listTasksGQL.watch({
      status: TaskStatus.Uncompleted,
      first: 100,
    });
  }

  public async ngOnInit(): Promise<void> {
    await this.initTasks();
  }

  private async initTasks(): Promise<void> {
    const pushIntoTasks = (_query: ListTasksQuery): void => {
      this.tasks.push(
        ..._query.tasks.edges
          // CompleteTask を実行した際、完了した Task 単体のキャッシュは更新される（status は TaskStatus.Completed になる）が、
          // 該当 Task が ListTasksQuery のキャッシュから除外されるわけではないことを考慮する必要がある。
          .filter((_edge) => _edge.node.status === TaskStatus.Uncompleted)
          .map((_edge) => _edge.node),
      );
    };

    let result: ApolloQueryResult<ListTasksQuery>;
    try {
      result = await this.listTasksQuery.result();
    } catch (e) {
      this.notificationService.notifyUnexpectedError(e);
      return;
    }

    pushIntoTasks(result.data);

    this.isTasksReady = true;

    while (result.data.tasks.pageInfo.hasNextPage) {
      try {
        result = await this.listTasksQuery.fetchMore({
          variables: {
            after: result.data.tasks.pageInfo.endCursor,
          },
        });
      } catch (e) {
        this.notificationService.notifyUnexpectedError(e);
        return;
      }

      pushIntoTasks(result.data);
    }
  }

  public async onChangeTaskCheckbox(task: Task): Promise<void> {
    if (!this.checkedTaskIDs.includes(task.id)) {
      return;
    }

    this.isTaskCompleting = true;

    try {
      await firstValueFrom(
        this.completeTaskGQL.mutate({
          id: task.id,
        }),
      );
    } catch (e) {
      this.notificationService.notifyUnexpectedError(e);
      this.isTaskCompleting = false;
      return;
    }

    await utils.sleep(500);

    this.tasks = this.tasks.filter((_task) => _task.id !== task.id);
    this.isTaskCompleting = false;
  }

  public onClickAddTaskButton(): void {
    // TODO
  }
}
