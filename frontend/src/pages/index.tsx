import { useEffect } from "react";

import { FieldMask } from "@bufbuild/protobuf";
import { Box, Container, Flex } from "@chakra-ui/react";
import { atom, useAtom } from "jotai";

import { CreateTaskForm, TaskRow } from "@/components";
import { TaskService } from "@/gen/app/v1/task_connect";
import {
  Task,
  TaskServiceCreateResponse,
  TaskServiceListResponse,
  TaskServiceUpdateResponse,
  TaskStatus,
} from "@/gen/app/v1/task_pb";
import { useConnectClient, useErrorToast, useSuccessToast } from "@/hooks";
import { CompleteTaskInputs, CreateTaskInputs } from "@/interfaces";
import { getErrorMessage } from "@/utils";

const tasksAtom = atom<Task[]>([]);

export default function HomePage() {
  const [tasks, setTasks] = useAtom(tasksAtom);

  const taskService = useConnectClient(TaskService);

  const errorToast = useErrorToast();
  const successToast = useSuccessToast();

  useEffect(() => {
    (async () => {
      let initialTasks: Task[];
      {
        let resp: TaskServiceListResponse;
        try {
          resp = await taskService.list({
            status: TaskStatus.UNCOMPLETED,
          });
        } catch (err) {
          errorToast({
            description: getErrorMessage(err),
          });
          return;
        }

        initialTasks = resp.tasks.reverse();
      }

      setTasks(initialTasks);
    })();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  async function onCreateTask(inputs: CreateTaskInputs): Promise<void> {
    let taskCreated: Task;
    {
      let resp: TaskServiceCreateResponse;
      try {
        resp = await taskService.create({
          title: inputs.title,
        });
      } catch (err) {
        errorToast({
          title: getErrorMessage(err),
        });
        return;
      }
      if (resp.task === undefined) {
        errorToast({
          title: "failed to create task",
        });
        return;
      }

      taskCreated = resp.task;
    }

    setTasks([taskCreated, ...tasks]);

    successToast({
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
        errorToast({
          description: getErrorMessage(err),
        });
        return;
      }
      if (resp.task === undefined) {
        errorToast({
          description: "failed to complete task",
        });
        return;
      }

      taskCompleted = resp.task;
    }

    setTasks(tasks.filter((task: Task) => task.id !== taskCompleted.id));

    successToast({
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
