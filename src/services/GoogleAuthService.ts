import { Common } from "googleapis";
import { GoogleTokens } from "../helpers/GoogleTokens";
import { PromiseHandler } from "../utils/PromiseHandler";
import { Either } from "../utils/Either";
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
    const clientToken = await PromiseHandler.wrap(this.client.getToken(code));
    return await Either.match(
      clientToken,
      (err) => {
        console.error("Failed", err);
        return {
          message: "Unauthorized",
        };
      },
      async (clientToken) => {
        const { tokens } = clientToken;
        this.client.setCredentials(tokens);
        return await this.saveAuthToken(tokens);
      },
    );
  }

  private async saveAuthToken(token: Credentials) {
    const savedToken = await PromiseHandler.wrap(
      this.GoogleTokens.saveTokens(token),
    );
    return await Either.match(
      savedToken,
      async (err) => {
        console.error("Failed", err);
        return { message: "Something went wrong" };
      },
      async () => ({ message: "Authorized" }),
    );
  }
}
