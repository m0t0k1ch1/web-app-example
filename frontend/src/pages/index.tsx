import { useEffect, useState } from "react";

import { Container, Flex } from "@chakra-ui/react";
import { useToast } from "@chakra-ui/react";

import { TaskForm, TaskRow } from "@/components";
import { TaskService } from "@/gen/app/v1/task_connect";
import { Task, TaskStatus } from "@/gen/app/v1/task_pb";
import { useConnectClient, useErrorToast, useSuccessToast } from "@/hooks";
import { TaskFormInputs } from "@/interfaces/TaskFormInputs";

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
  }, []);

  async function onSubmit(inputs: TaskFormInputs): Promise<void> {
    let newTask: Task;
    {
      const resp = await taskService.create({
        title: inputs.title,
      });
      if (resp.task === undefined) {
        eToast({
          description: "Failed to create task",
        });
        return;
      }

      newTask = resp.task;
    }

    setTasks([newTask, ...tasks]);

    sToast({
      title: "Task created",
    });
  }

  return (
    <Container h="100%">
      <Flex direction="column">
        <TaskForm onSubmit={onSubmit} />
        <Flex direction="column">
          {tasks.map((task: Task) => (
            <TaskRow key={task.id} task={task} />
          ))}
        </Flex>
      </Flex>
    </Container>
  );
}
