import { Hono } from "hono";

interface RouterContract {
  getApp(): Hono;
}

export abstract class BaseRouter implements RouterContract {
  protected app: Hono;

  constructor(app: Hono) {
    this.app = app;
  }

  protected abstract routes(): void;

  getApp(): Hono {
    return this.app;
  }
}
