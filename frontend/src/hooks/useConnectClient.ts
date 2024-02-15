import { useMemo } from "react";

import { ServiceType } from "@bufbuild/protobuf";
import { createPromiseClient, PromiseClient } from "@connectrpc/connect";
import { createConnectTransport } from "@connectrpc/connect-web";

const transport = createConnectTransport({
  baseUrl: process.env.NEXT_PUBLIC_APP_BACKEND_BASE_URL,
});

export function useConnectClient<T extends ServiceType>(
  service: T
): PromiseClient<T> {
  return useMemo(() => createPromiseClient(service, transport), [service]);
}
