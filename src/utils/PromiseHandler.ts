type PromiseHandlerValue<T> = { ok: T; err: null } | { ok: null; err: Error };
type TryPromiseParams<T, R> = {
  try: (promise: T) => R | Promise<R>;
  catch: (err: Error) => Error | R;
};

export class PromiseHandler {
  private static async wrap<T>(
    promise: Promise<T>,
  ): Promise<PromiseHandlerValue<T>> {
    try {
      const ok = await promise;
      return { ok, err: null };
    } catch (err) {
      if (err instanceof Error) return { ok: null, err };
      return { ok: null, err: new Error("Unknown error") };
    }
  }

  public static async tryPromise<T, R>(
    promise: Promise<T>,
    tryPromise: TryPromiseParams<T, R>,
  ): Promise<R | Error> {
    const { ok, err } = await this.wrap(promise);
    if (err !== null) {
      return tryPromise.catch(err);
    }
    if (ok !== null) {
      return tryPromise.try(ok);
    }
    return tryPromise.catch(new Error("Unexpected error"));
  }
}
