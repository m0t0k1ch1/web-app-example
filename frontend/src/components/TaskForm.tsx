import {
  Button,
  Flex,
  FormControl,
  FormErrorMessage,
  Input,
} from "@chakra-ui/react";
import { useForm, SubmitHandler } from "react-hook-form";

import { TaskFormInputs } from "@/interfaces/TaskFormInputs";

interface Props {
  onSubmit: SubmitHandler<TaskFormInputs>;
}

export function TaskForm({ onSubmit }: Props) {
  const {
    register,
    handleSubmit,
    formState: { errors, isSubmitting },
  } = useForm<TaskFormInputs>();

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <Flex gap="2">
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
        <Button type="submit" isLoading={isSubmitting}>
          Add
        </Button>
      </Flex>
    </form>
  );
}
