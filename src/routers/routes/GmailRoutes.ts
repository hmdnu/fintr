import { Hono } from "hono";
import { BaseRouter } from "../base/BaseRouter";
import { GmailService } from "../../services/GmailService";

export class GmailRoutes extends BaseRouter {
  private GmailService: GmailService;

  constructor(GmailService: GmailService) {
    super(new Hono());
    this.GmailService = GmailService;
    this.routes();
  }

  protected routes(): void {
    this.app.get("/", (c) => c.text("/gmail"));
  }
}
