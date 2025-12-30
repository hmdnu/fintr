import { Service, ServiceRegistry } from "../services/ServiceRegistry";
import { GmailRoutes } from "./routes/GmailRoutes";
import { GoogleAuthRoutes } from "./routes/GoogleAuthRoutes";

export type Routes = {
  googleAuthRoutes: GoogleAuthRoutes;
  gmailRoutes: GmailRoutes;
};

export class RoutesRegistry {
  private service: Service;

  constructor(ServiceRegistry: ServiceRegistry) {
    this.service = ServiceRegistry.instantiate();
  }

  public instantiate(): Routes {
    const googleAuthRoutes = new GoogleAuthRoutes(this.service.googleAuthService);
    const gmailRoutes = new GmailRoutes(this.service.gmailService);

    return { googleAuthRoutes, gmailRoutes };
  }
}
