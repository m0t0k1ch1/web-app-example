import { Text } from "@chakra-ui/react";

import { Task } from "@/gen/app/v1/task_pb";

interface Props {
  task: Task;
}

export function TaskRow({ task }: Props) {
  return <Text>{task.title}</Text>;
}
