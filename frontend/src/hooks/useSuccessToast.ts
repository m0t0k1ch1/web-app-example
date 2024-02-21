import {
  useToast,
  CreateToastFnReturn,
  UseToastOptions,
} from "@chakra-ui/react";

const DEFAULT_OPTIONS: UseToastOptions = {
  duration: 3_000,
  title: "Success",
  isClosable: true,
  status: "success",
} as const;

export function useSuccessToast(
  options: UseToastOptions = {}
): CreateToastFnReturn {
  return useToast({
    ...DEFAULT_OPTIONS,
    ...options,
  });
}
