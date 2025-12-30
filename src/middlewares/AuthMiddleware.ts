import { createMiddleware } from "hono/factory";
import { GoogleTokens } from "../helpers/GoogleTokens";

export class AuthMiddleware {
  private googleTokens: GoogleTokens;

  constructor(googleTokens: GoogleTokens) {
    this.googleTokens = googleTokens;
  }

  public authorize() {
    return createMiddleware(async (c, next) => {
      const token = await this.googleTokens.loadTokens();
      if (token instanceof Error) {
        c.status(401);
        return c.json({ message: "Unauthorized", reason: token.message });
      }

      await next();
    });
  }
}
