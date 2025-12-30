import { Hono } from "hono";

export abstract class BaseRouter {
  protected app: Hono;

  constructor(app: Hono) {
    this.app = app;
  }

  protected abstract routes(): void;

  getApp(): Hono {
    return this.app;
  }
}
