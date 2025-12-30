import { Hono } from "hono";
import { BaseRouter } from "../base/BaseRouter";
import { GmailService } from "../../services/GmailService";
import { AuthMiddleware } from "../../middlewares/AuthMiddleware";

type GmailRoutesType = {
  gmailService: GmailService;
  authMiddleware: AuthMiddleware;
};

export class GmailRoutes extends BaseRouter {
  private gmailService: GmailService;
  private authMiddleware: AuthMiddleware;

  constructor({ gmailService, authMiddleware }: GmailRoutesType) {
    super(new Hono());
    this.gmailService = gmailService;
    this.authMiddleware = authMiddleware;
    this.routes();
  }

  protected routes(): void {
    this.app.get("/list", this.authMiddleware.authorize(), (c) => {
      return c.text("/gmail");
    });
  }
}
