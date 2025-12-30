import { google } from "googleapis";
import { GoogleTokens } from "../helpers/GoogleTokens";

export class GmailService {
  private GoogleTokens: GoogleTokens;

  constructor(googleTokens: GoogleTokens) {
    this.GoogleTokens = googleTokens;
  }

  public async listMails() {
    const token = await this.GoogleTokens.loadTokens();
    const client = this.GoogleTokens.getOauthClient();
  }
}
