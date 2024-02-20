import { ChangeEvent } from "react";

import { Box, Checkbox, Flex, Text } from "@chakra-ui/react";
import { atom, useAtom } from "jotai";

import { Task } from "@/gen/app/v1/task_pb";
import { CompleteTaskInputs } from "@/interfaces";
import { sleep } from "@/utils";

const isSubmittingAtom = atom<boolean>(false);

interface Props {
  task: Task;
  onComplete: (inputs: CompleteTaskInputs) => Promise<void>;
}

export function TaskRow(props: Props) {
  const [isSubmitting, setIsSubmitting] = useAtom(isSubmittingAtom);

  async function onChange(event: ChangeEvent<HTMLInputElement>): Promise<void> {
    if (!event.target.checked) {
      return;
    }

    setIsSubmitting(true);

    await sleep(500);

    await props.onComplete({
      id: props.task.id,
    });

    setIsSubmitting(false);
  }

  return (
    <Box border="1px" borderColor="gray.200" borderRadius={6} p={4}>
      <Flex alignItems="start" direction="row" gap={4}>
        <Checkbox
          colorScheme="orange"
          h={6}
          isReadOnly={isSubmitting}
          onChange={onChange}
          size="lg"
        />
        <Text lineHeight={6}>{props.task.title}</Text>
      </Flex>
    </Box>
  );
}
