export function getErrorMessage(err: unknown): string {
  if (err instanceof Error) {
    return err.message;
  }

  console.error(err);

  return "unexpected error occured";
}

export async function sleep(msec: number): Promise<void> {
  return new Promise<void>((resolve) => setTimeout(resolve, msec));
}
