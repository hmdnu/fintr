import { Hono } from "hono";
import { Routes, RoutesRegistry } from "./RouterRegistry";

export class Router {
  private app = new Hono();
  private routeRegistry: Routes;

  constructor(RouteRegistry: RoutesRegistry) {
    this.routeRegistry = RouteRegistry.instantiate();
    this.bindRoutes();
  }

  private bindRoutes() {
    this.app.get("/", (c) => {
      return c.json({ name: "Fintr (Financial Tracker)", version: "0.1.0" });
    });
    this.app.route("/auth", this.routeRegistry.googleAuthRoutes.getApp());
    this.app.route("/gmail", this.routeRegistry.gmailRoutes.getApp());
  }

  public getApp() {
    return this.app;
  }
}
