import {
  Button,
  Flex,
  FormControl,
  FormErrorMessage,
  Input,
} from "@chakra-ui/react";
import { useForm } from "react-hook-form";

import { CreateTaskInputs } from "@/interfaces";

interface Props {
  onSubmit: (inputs: CreateTaskInputs) => Promise<void>;
}

export function CreateTaskForm(props: Props) {
  const {
    formState: { errors, isSubmitting },
    handleSubmit,
    register,
    reset,
  } = useForm<CreateTaskInputs>();

  async function onSubmit(inputs: CreateTaskInputs): Promise<void> {
    await props.onSubmit(inputs);

    reset();
  }

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <Flex direction="row" gap="2">
        <FormControl isInvalid={errors.title !== undefined}>
          <Input
            placeholder="title"
            {...register("title", {
              required: "required",
              maxLength: {
                value: 32,
                message: "must be 32 characters or less",
              },
            })}
          />
          <FormErrorMessage>
            {errors.title !== undefined && errors.title.message}
          </FormErrorMessage>
        </FormControl>
        <Button colorScheme="teal" isLoading={isSubmitting} type="submit">
          Add
        </Button>
      </Flex>
    </form>
  );
}
