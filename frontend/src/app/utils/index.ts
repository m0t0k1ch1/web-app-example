export async function sleep(msec: number): Promise<void> {
  return new Promise((resolve) => setTimeout(resolve, msec));
}

export function stringifyError(err: any): string {
  if (typeof err === 'string') {
    return err;
  }
  if ('message' in err) {
    return err.message;
  }

  throw err;
}
