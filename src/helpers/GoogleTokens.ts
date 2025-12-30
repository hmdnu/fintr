import { google } from "googleapis";
import type { Credentials } from "google-auth-library";
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
  private TOKEN_PATH = "src/credentials/token.json";

  constructor(googleClientConfig: GoogleClientConfig) {
    this.CLIENT_ID = googleClientConfig.clientId;
    this.CLIENT_SECRET = googleClientConfig.clientSecret;
    this.REDIRECT_URI = googleClientConfig.redirectUri;
  }

  public async loadTokens() {
    return await PromiseHandler.tryPromise(Bun.file(this.TOKEN_PATH).text(), {
      try: (token) => token,
      catch: (err) => {
        return err;
      },
    });
  }

  public async saveTokens(tokens: Credentials) {
    return await PromiseHandler.tryPromise(
      Bun.write(this.TOKEN_PATH, JSON.stringify(tokens)),
      { try: () => null, catch: (err) => err },
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
