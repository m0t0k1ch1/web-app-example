import { createPromiseClient, PromiseClient } from "@connectrpc/connect";
import { createConnectTransport } from "@connectrpc/connect-web";

import { TaskService } from "@/gen/app/v1/task_connect";

export class Backend {
  public readonly taskService: PromiseClient<typeof TaskService>;

  constructor() {
    const transport = createConnectTransport({
      baseUrl: process.env.APP_BACKEND_BASE_URL,
    });

    this.taskService = createPromiseClient(TaskService, transport);
  }
}

export const backend = new Backend();
