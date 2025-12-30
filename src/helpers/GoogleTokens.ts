import { google } from "googleapis";
import type { Credentials } from "google-auth-library";
import { Either } from "../utils/Either";
import { PromiseHandler } from "../utils/PromiseHandler";

type GoogleClientConfig = {
  clientId: string;
  clientSecret: string;
  redirectUri: string;
};

export class GoogleTokens {
  private CLIENT_ID = "";
  private CLIENT_SECRET = "";
  private REDIRECT_URI = "";
  private TOKEN_PATH = "../credentials/token.json";

  constructor(googleClientConfig: GoogleClientConfig) {
    this.CLIENT_ID = googleClientConfig.clientId;
    this.CLIENT_SECRET = googleClientConfig.clientSecret;
    this.REDIRECT_URI = googleClientConfig.redirectUri;
  }

  public loadTokens() {
    return Bun.file(this.TOKEN_PATH).text();
  }

  public async saveTokens(tokens: Credentials) {
    const writtenFile = await PromiseHandler.wrap(
      Bun.write(this.TOKEN_PATH, JSON.stringify(tokens)),
    );

    return Either.match(
      writtenFile,
      async (err) => err,
      () => null,
    );
  }

  public getOauthClient() {
    return new google.auth.OAuth2(
      this.CLIENT_ID,
      this.CLIENT_SECRET,
      this.REDIRECT_URI,
    );
  }
}
