import { GoogleTokens } from "../helpers/GoogleTokens";
import { GmailService } from "./GmailService";
import { GoogleAuthService } from "./GoogleAuthService";

export type Service = {
  googleAuthService: GoogleAuthService;
  gmailService: GmailService;
};

export class ServiceRegistry {
  private GoogleTokens: GoogleTokens;

  constructor(googleToken: GoogleTokens) {
    this.GoogleTokens = googleToken;
  }

  public instantiate(): Service {
    const googleAuthService = new GoogleAuthService(this.GoogleTokens);
    const gmailService = new GmailService(this.GoogleTokens);

    return { googleAuthService, gmailService };
  }
}
