export function stringifyError(err: any): string {
  if (typeof err === 'string') {
    return err;
  }
  if (err instanceof Error) {
    return err.message;
  }

  throw err;
}
