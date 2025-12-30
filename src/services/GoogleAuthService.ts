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
    return await PromiseHandler.tryPromise(this.client.getToken(code), {
      try: async ({ tokens }) => {
        this.client.setCredentials(tokens);
        return this.saveAuthToken(tokens);
      },
      catch: () => {
        return { message: "Unauthorized", status: 401 };
      },
    });
  }

  private async saveAuthToken(token: Credentials) {
    return await PromiseHandler.tryPromise(
      this.GoogleTokens.saveTokens(token),
      {
        try: () => {
          return { message: "Authorized", status: 200 };
        },
        catch: (err) => {
          console.error("Failed", err);
          return { message: "Unauthorized", status: 401 };
        },
      },
    );
  }
}
