import {
  useToast,
  CreateToastFnReturn,
  UseToastOptions,
} from "@chakra-ui/react";

export function useErrorToast(
  options: UseToastOptions = {}
): CreateToastFnReturn {
  return useToast({
    title: "Error",
    status: "error",
    duration: 5_000,
    isClosable: true,
    ...options,
  });
}
