import {
  useToast,
  CreateToastFnReturn,
  UseToastOptions,
} from "@chakra-ui/react";

export function useSuccessToast(
  options: UseToastOptions = {}
): CreateToastFnReturn {
  return useToast({
    title: "Success",
    status: "success",
    duration: 3_000,
    isClosable: true,
    ...options,
  });
}
