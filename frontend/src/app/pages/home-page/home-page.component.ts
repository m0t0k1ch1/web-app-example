import { Component, OnInit, inject } from '@angular/core';
import { CommonModule } from '@angular/common';
import { Router } from '@angular/router';
import { MatRippleModule } from '@angular/material/core';

import { ApolloQueryResult } from '@apollo/client/core';
import { QueryRef } from 'apollo-angular';

import {
  ListTasksGQL,
  ListTasksQuery,
  ListTasksQueryVariables,
  Task,
  TaskStatus,
} from '../../../gen/graphql-codegen/schema';

import { NotificationService } from '../../services/notification.service';

@Component({
  selector: 'app-home-page',
  standalone: true,
  imports: [CommonModule, MatRippleModule],
  templateUrl: './home-page.component.html',
  styleUrl: './home-page.component.css',
})
export class HomePageComponent implements OnInit {
  private router = inject(Router);

  private listTasksGQL = inject(ListTasksGQL);

  private notificationService = inject(NotificationService);

  private listTasksQuery: QueryRef<ListTasksQuery, ListTasksQueryVariables>;

  public tasks: Task[] = [];
  public isTasksInitialized: boolean = false;

  constructor() {
    this.listTasksQuery = this.listTasksGQL.watch({
      status: TaskStatus.Uncompleted,
      first: 100,
    });
  }

  async ngOnInit(): Promise<void> {
    await this.initTasks();
  }

  private async initTasks() {
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

  public onClickAddTaskButton(): void {
    this.router.navigate(['hogepi']);
  }
}
