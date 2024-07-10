import { Component, OnInit, inject } from '@angular/core';
import { CommonModule } from '@angular/common';
import { MatRippleModule } from '@angular/material/core';
import { firstValueFrom } from 'rxjs';

import { ApolloQueryResult } from '@apollo/client/core';
import { QueryRef } from 'apollo-angular';

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
  imports: [CommonModule, MatRippleModule],
  templateUrl: './home-page.component.html',
  styleUrl: './home-page.component.css',
})
export class HomePageComponent implements OnInit {
  private listTasksGQL = inject(ListTasksGQL);
  private completeTaskGQL = inject(CompleteTaskGQL);

  private notificationService = inject(NotificationService);

  private listTasksQuery: QueryRef<ListTasksQuery, ListTasksQueryVariables>;

  public tasks: Task[] = [];
  public isTasksInitialized: boolean = false;
  public isTaskCompleting: boolean = false;

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
    let result: ApolloQueryResult<ListTasksQuery>;
    try {
      result = await this.listTasksQuery.result();
    } catch (e) {
      this.notificationService.notifyUnexpectedError(e);
      return;
    }

    this.tasks.push(...result.data.tasks.edges.map((edge) => edge.node));
    this.isTasksInitialized = true;

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

      this.tasks.push(...result.data.tasks.edges.map((edge) => edge.node));
    }
  }

  public async onChangeTaskStatus(task: Task, event: Event): Promise<void> {
    const isChecked = (event.target as HTMLInputElement).checked;
    if (!isChecked) {
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
