import { Either, EitherT } from "./Either";

export class PromiseHandler {
  public static async wrap<T>(promise: Promise<T>): Promise<EitherT<Error, T>> {
    try {
      const result = await promise;
      return Either.right(result);
    } catch (err) {
      if (err instanceof Error) return Either.left(err);
      return Either.left(new Error("Unknown error"));
    }
  }
}
