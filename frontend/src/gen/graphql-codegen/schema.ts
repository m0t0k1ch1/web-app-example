import { gql } from 'apollo-angular';
import { Injectable } from '@angular/core';
import * as Apollo from 'apollo-angular';
export type Maybe<T> = T | null;
export type InputMaybe<T> = Maybe<T>;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
export type MakeEmpty<T extends { [key: string]: unknown }, K extends keyof T> = { [_ in K]?: never };
export type Incremental<T> = T | { [P in keyof T]?: P extends ' $fragmentName' | '__typename' ? T[P] : never };
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: { input: string; output: string; }
  String: { input: string; output: string; }
  Boolean: { input: boolean; output: boolean; }
  Int: { input: number; output: number; }
  Float: { input: number; output: number; }
  Int32: { input: number; output: number; }
  Int64: { input: number; output: number; }
  Uint: { input: number; output: number; }
  Uint32: { input: number; output: number; }
  Uint64: { input: number; output: number; }
};

export type CompleteTaskInput = {
  clientMutationId?: InputMaybe<Scalars['String']['input']>;
  id: Scalars['ID']['input'];
};

export type CompleteTaskPayload = {
  __typename?: 'CompleteTaskPayload';
  clientMutationId?: Maybe<Scalars['String']['output']>;
  task?: Maybe<Task>;
};

export type CreateTaskInput = {
  clientMutationId?: InputMaybe<Scalars['String']['input']>;
  title: Scalars['String']['input'];
};

export type CreateTaskPayload = {
  __typename?: 'CreateTaskPayload';
  clientMutationId?: Maybe<Scalars['String']['output']>;
  task: Task;
};

export type Mutation = {
  __typename?: 'Mutation';
  completeTask: CompleteTaskPayload;
  createTask: CreateTaskPayload;
  noop: NoopPayload;
};


export type MutationCompleteTaskArgs = {
  input: CompleteTaskInput;
};


export type MutationCreateTaskArgs = {
  input: CreateTaskInput;
};


export type MutationNoopArgs = {
  input: NoopInput;
};

export type Node = {
  id: Scalars['ID']['output'];
};

export type NoopInput = {
  clientMutationId?: InputMaybe<Scalars['String']['input']>;
};

export type NoopPayload = {
  __typename?: 'NoopPayload';
  clientMutationId?: Maybe<Scalars['String']['output']>;
};

export type PageInfo = {
  __typename?: 'PageInfo';
  endCursor?: Maybe<Scalars['String']['output']>;
  hasNextPage: Scalars['Boolean']['output'];
};

export type Query = {
  __typename?: 'Query';
  node?: Maybe<Node>;
  tasks: TaskConnection;
};


export type QueryNodeArgs = {
  id: Scalars['ID']['input'];
};


export type QueryTasksArgs = {
  after?: InputMaybe<Scalars['String']['input']>;
  first: Scalars['Int32']['input'];
  status?: InputMaybe<TaskStatus>;
};

export type Task = Node & {
  __typename?: 'Task';
  id: Scalars['ID']['output'];
  status: TaskStatus;
  title: Scalars['String']['output'];
};

export type TaskConnection = {
  __typename?: 'TaskConnection';
  edges: Array<TaskEdge>;
  nodes: Array<Task>;
  pageInfo: PageInfo;
  totalCount: Scalars['Int64']['output'];
};

export type TaskEdge = {
  __typename?: 'TaskEdge';
  cursor: Scalars['String']['output'];
  node: Task;
};

export enum TaskStatus {
  Completed = 'COMPLETED',
  Uncompleted = 'UNCOMPLETED'
}

export type CreateTaskMutationVariables = Exact<{
  title: Scalars['String']['input'];
}>;


export type CreateTaskMutation = { __typename?: 'Mutation', createTask: { __typename?: 'CreateTaskPayload', task: { __typename?: 'Task', id: string, title: string, status: TaskStatus } } };

export type CompleteTaskMutationVariables = Exact<{
  id: Scalars['ID']['input'];
}>;


export type CompleteTaskMutation = { __typename?: 'Mutation', completeTask: { __typename?: 'CompleteTaskPayload', task?: { __typename?: 'Task', id: string, title: string, status: TaskStatus } | null } };

export type ListTasksQueryVariables = Exact<{
  status: TaskStatus;
  after?: InputMaybe<Scalars['String']['input']>;
  first: Scalars['Int32']['input'];
}>;


export type ListTasksQuery = { __typename?: 'Query', tasks: { __typename?: 'TaskConnection', edges: Array<{ __typename?: 'TaskEdge', node: { __typename?: 'Task', id: string, title: string, status: TaskStatus } }>, pageInfo: { __typename?: 'PageInfo', endCursor?: string | null, hasNextPage: boolean } } };

export const CreateTaskDocument = gql`
    mutation CreateTask($title: String!) {
  createTask(input: {title: $title}) {
    task {
      id
      title
      status
    }
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class CreateTaskGQL extends Apollo.Mutation<CreateTaskMutation, CreateTaskMutationVariables> {
    document = CreateTaskDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const CompleteTaskDocument = gql`
    mutation CompleteTask($id: ID!) {
  completeTask(input: {id: $id}) {
    task {
      id
      title
      status
    }
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class CompleteTaskGQL extends Apollo.Mutation<CompleteTaskMutation, CompleteTaskMutationVariables> {
    document = CompleteTaskDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const ListTasksDocument = gql`
    query ListTasks($status: TaskStatus!, $after: String, $first: Int32!) {
  tasks(status: $status, after: $after, first: $first) {
    edges {
      node {
        id
        title
        status
      }
    }
    pageInfo {
      endCursor
      hasNextPage
    }
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class ListTasksGQL extends Apollo.Query<ListTasksQuery, ListTasksQueryVariables> {
    document = ListTasksDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }