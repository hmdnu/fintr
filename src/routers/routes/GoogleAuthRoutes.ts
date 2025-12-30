import { Hono } from "hono";
import { GoogleAuthService } from "../../services/GoogleAuthService";
import { BaseRouter } from "../base/BaseRouter";

export class GoogleAuthRoutes extends BaseRouter {
  private GoogleAuthService: GoogleAuthService;

  constructor(GoogleAuthService: GoogleAuthService) {
    super(new Hono());
    this.GoogleAuthService = GoogleAuthService;
    this.routes();
  }

  protected routes(): void {
    this.app.get("/", (c) => {
      const url = this.GoogleAuthService.generateAuthUrl();
      return c.redirect(url);
    });

    this.app.get("/oauth2callback", async (c) => {
      const { code } = c.req.query();
      const oauthResponse = await this.GoogleAuthService.oauthCallback(code);
      return c.json(oauthResponse);
    });
  }
}
