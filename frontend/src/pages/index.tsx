import { GetServerSideProps, InferGetServerSidePropsType } from "next";

import { Container } from "@chakra-ui/react";

import { TaskForm } from "@/components";
import { Task } from "@/gen/app/v1/task_pb";
import { TaskFormInputs } from "@/interfaces/TaskFormInputs";
import { backend } from "@/lib/backend";

interface Props {
  tasks: Task[];
}

export const getServerSideProps: GetServerSideProps<Props> = async () => {
  const { tasks } = await backend.taskService.list({});

  return { props: { tasks } };
};

export default function Page({
  tasks,
}: InferGetServerSidePropsType<typeof getServerSideProps>) {
  const onSubmit = async (inputs: TaskFormInputs): Promise<void> => {
    console.log(inputs.title);
  };

  return (
    <Container h="100%">
      <TaskForm onSubmit={onSubmit} />
    </Container>
  );
}
