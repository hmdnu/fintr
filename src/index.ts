import { GoogleTokens } from "./helpers/GoogleTokens";
import { Router } from "./routers/Router";
import { RoutesRegistry } from "./routers/RouterRegistry";
import { ServiceRegistry } from "./services/ServiceRegistry";

const googleTokens = new GoogleTokens({
  clientId: String(Bun.env.GOOGLE_CLIENT_ID),
  clientSecret: String(Bun.env.GOOGLE_CLIENT_SECRET),
  redirectUri: String(Bun.env.GOOGLE_REDIRECT_URI),
});
const serviceRegistry = new ServiceRegistry(googleTokens);
const routesRegistry = new RoutesRegistry(serviceRegistry);
const router = new Router(routesRegistry);

export default router.getApp();
