import { useEffect, useState } from "react";

import { FieldMask } from "@bufbuild/protobuf";
import { Box, Container, Flex } from "@chakra-ui/react";

import { CreateTaskForm, TaskRow } from "@/components";
import { TaskService } from "@/gen/app/v1/task_connect";
import {
  Task,
  TaskServiceCreateResponse,
  TaskServiceUpdateResponse,
  TaskStatus,
} from "@/gen/app/v1/task_pb";
import { useConnectClient, useErrorToast, useSuccessToast } from "@/hooks";
import { CompleteTaskInputs, CreateTaskInputs } from "@/interfaces";
import { getErrorMessage } from "@/utils";

export default function Page() {
  const [tasks, setTasks] = useState<Task[]>([]);

  const taskService = useConnectClient(TaskService);

  const eToast = useErrorToast();
  const sToast = useSuccessToast();

  useEffect(() => {
    (async () => {
      let initialTasks: Task[];
      {
        const resp = await taskService.list({
          status: TaskStatus.UNCOMPLETED,
        });

        initialTasks = resp.tasks.reverse();
      }

      setTasks(initialTasks);
    })();
  }, [taskService]);

  async function onCreateTask(inputs: CreateTaskInputs): Promise<void> {
    let taskCreated: Task;
    {
      let resp: TaskServiceCreateResponse;
      try {
        resp = await taskService.create({
          title: inputs.title,
        });
      } catch (err) {
        eToast({
          description: getErrorMessage(err),
        });
        return;
      }
      if (resp.task === undefined) {
        eToast({
          description: "failed to create task",
        });
        return;
      }

      taskCreated = resp.task;
    }

    setTasks([taskCreated, ...tasks]);

    sToast({
      title: "task created",
    });
  }

  async function onCompleteTask(inputs: CompleteTaskInputs): Promise<void> {
    let taskCompleted: Task;
    {
      let resp: TaskServiceUpdateResponse;
      try {
        resp = await taskService.update({
          task: {
            id: inputs.id,
            status: TaskStatus.COMPLETED,
          },
          fieldMask: new FieldMask({
            paths: ["id", "status"],
          }),
        });
      } catch (err) {
        eToast({
          description: getErrorMessage(err),
        });
        return;
      }
      if (resp.task === undefined) {
        eToast({
          description: "failed to complete task",
        });
        return;
      }

      taskCompleted = resp.task;
    }

    setTasks(tasks.filter((task: Task) => task.id !== taskCompleted.id));

    sToast({
      title: "task completed",
    });
  }

  return (
    <Container my={4}>
      <Flex direction="column">
        <Box h={20}>
          <CreateTaskForm onSubmit={onCreateTask} />
        </Box>
        <Flex direction="column" gap={2}>
          {tasks.map((task: Task) => (
            <TaskRow key={task.id} onComplete={onCompleteTask} task={task} />
          ))}
        </Flex>
      </Flex>
    </Container>
  );
}
