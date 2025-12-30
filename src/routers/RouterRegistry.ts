import { AuthMiddleware } from "../middlewares/AuthMiddleware";
import { Service, ServiceRegistry } from "../services/ServiceRegistry";
import { GmailRoutes } from "./routes/GmailRoutes";
import { GoogleAuthRoutes } from "./routes/GoogleAuthRoutes";

export type Routes = {
  googleAuthRoutes: GoogleAuthRoutes;
  gmailRoutes: GmailRoutes;
};

export class RoutesRegistry {
  private service: Service;
  private authMiddleware: AuthMiddleware;

  constructor(
    authMiddleware: AuthMiddleware,
    serviceRegistry: ServiceRegistry,
  ) {
    this.service = serviceRegistry.instantiate();
    this.authMiddleware = authMiddleware;
  }

  public instantiate(): Routes {
    const googleAuthRoutes = new GoogleAuthRoutes(
      this.service.googleAuthService,
    );
    const gmailRoutes = new GmailRoutes({
      gmailService: this.service.gmailService,
      authMiddleware: this.authMiddleware,
    });

    return { googleAuthRoutes, gmailRoutes };
  }
}
