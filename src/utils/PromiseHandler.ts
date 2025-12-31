export type Result<T, E extends Error = Error> =
  | { ok: true; value: T }
  | { ok: false; err: E };

export class PromiseHandler {
  public static async wrap<T>(promise: Promise<T>): Promise<Result<T>> {
    try {
      return { ok: true, value: await promise };
    } catch (e: unknown) {
      return {
        ok: false,
        err: e instanceof Error ? e : new Error(`Unknown error: ${e}`),
      };
    }
  }
}
