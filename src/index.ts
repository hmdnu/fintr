import { GoogleTokens } from "./helpers/GoogleTokens";
import { AuthMiddleware } from "./middlewares/AuthMiddleware";
import { Router } from "./routers/Router";
import { RoutesRegistry } from "./routers/RouterRegistry";
import { ServiceRegistry } from "./services/ServiceRegistry";

const googleTokens = new GoogleTokens({
  clientId: String(Bun.env.GOOGLE_CLIENT_ID),
  clientSecret: String(Bun.env.GOOGLE_CLIENT_SECRET),
  redirectUri: String(Bun.env.GOOGLE_REDIRECT_URI),
});
const authMiddleware = new AuthMiddleware(googleTokens);
const serviceRegistry = new ServiceRegistry(googleTokens);
const routesRegistry = new RoutesRegistry(authMiddleware, serviceRegistry);
const router = new Router(routesRegistry);

export default router.getApp();
