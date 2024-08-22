import { TaskStatus } from '../../../../gen/graphql-codegen/schema';

export interface Task {
  id: string;
  title: string;
  status: TaskStatus;
}
