import { Component, OnInit, inject } from '@angular/core';
import { CommonModule } from '@angular/common';
import { Router } from '@angular/router';
import { MatRippleModule } from '@angular/material/core';

import { ApolloQueryResult } from '@apollo/client/core';
import { Apollo, QueryRef, gql } from 'apollo-angular';

import { NotificationService } from '../../services/notification.service';

const LIST_TASKS_QUERY = gql`
  query ListTasks($status: TaskStatus!, $after: String, $first: Int32!) {
    tasks(status: $status, after: $after, first: $first) {
      edges {
        node {
          id
          title
        }
      }
      pageInfo {
        endCursor
        hasNextPage
      }
    }
  }
`;

@Component({
  selector: 'app-home-page',
  standalone: true,
  imports: [CommonModule, MatRippleModule],
  templateUrl: './home-page.component.html',
  styleUrl: './home-page.component.css',
})
export class HomePageComponent implements OnInit {
  private router = inject(Router);
  private apollo = inject(Apollo);

  private notificationService = inject(NotificationService);

  private listTasksQuery: QueryRef<any, any>;

  public tasks: any[] = [];
  public isTasksInitialized: boolean = false;

  constructor() {
    this.listTasksQuery = this.apollo.watchQuery({
      query: LIST_TASKS_QUERY,
      variables: {
        status: 'UNCOMPLETED',
        first: 1,
      },
    });
  }

  async ngOnInit(): Promise<void> {
    await this.initTasks();
  }

  private async initTasks() {
    let result: ApolloQueryResult<any>;
    try {
      result = await this.listTasksQuery.result();
    } catch (e) {
      this.notificationService.notifyUnexpectedError(e);
      return;
    }

    this.tasks.push(...result.data.tasks.edges.map((edge: any) => edge.node));
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

      this.tasks.push(...result.data.tasks.edges.map((edge: any) => edge.node));
    }
  }

  public onClickAddTaskButton(): void {
    this.router.navigate(['hogepi']);
  }
}
