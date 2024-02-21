import {
  useToast,
  CreateToastFnReturn,
  UseToastOptions,
} from "@chakra-ui/react";

const DEFAULT_OPTIONS: UseToastOptions = {
  duration: 5_000,
  title: "Error",
  isClosable: true,
  status: "error",
} as const;

export function useErrorToast(
  options: UseToastOptions = {}
): CreateToastFnReturn {
  return useToast({
    ...DEFAULT_OPTIONS,
    ...options,
  });
}
