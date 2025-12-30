export type EitherT<L, R> =
  | { tag: "Left"; left: L }
  | { tag: "Right"; right: R };

export class Either {
  public static left<L, R = never>(left: L): EitherT<L, R> {
    return { tag: "Left", left };
  }

  public static right<R, L = never>(right: R): EitherT<L, R> {
    return { tag: "Right", right };
  }

  public static async match<L, R, T>(
    value: EitherT<L, R>,
    onLeft: (l: L) => T | Promise<T>,
    onRight: (r: R) => T | Promise<T>,
  ): Promise<T> {
    const resolved = value;
    return resolved.tag === "Left"
      ? await onLeft(resolved.left)
      : await onRight(resolved.right);
  }

  public static isRight<L, R>(
    e: EitherT<L, R>,
  ): e is { tag: "Right"; right: R } {
    return e.tag === "Right";
  }
}
