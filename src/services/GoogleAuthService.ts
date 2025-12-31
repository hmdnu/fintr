import { Common } from "googleapis";
import { GoogleTokens } from "../helpers/GoogleTokens";
import { PromiseHandler } from "../utils/PromiseHandler";
import { Credentials } from "google-auth-library";

export class GoogleAuthService {
  private GoogleTokens: GoogleTokens;
  private client: Common.OAuth2Client;
  private scopes = ["https://www.googleapis.com/auth/gmail.readonly"];

  constructor(GoogleTokens: GoogleTokens) {
    this.GoogleTokens = GoogleTokens;
    this.client = this.GoogleTokens.getOauthClient();
  }

  public generateAuthUrl() {
    return this.client.generateAuthUrl({
      access_type: "offline",
      prompt: "consent",
      scope: this.scopes,
    });
  }

  public async oauthCallback(code: string) {
    const token = await PromiseHandler.wrap(this.client.getToken(code));
    if (!token.ok) return { message: "Unauthorized", status: 401 };

    this.client.setCredentials(token.value.tokens);
    const savedToken = await PromiseHandler.wrap(
      this.GoogleTokens.saveTokens(token.value.tokens),
    );

    if (!savedToken.ok) return { message: "Unauthorized", status: 401 };
    return { message: "Authorized", status: 200 };
  }
}
