import { google } from "googleapis";
import { GoogleTokens } from "../helpers/GoogleTokens";

export class GmailService {
  private GoogleTokens: GoogleTokens;

  constructor(googleTokens: GoogleTokens) {
    this.GoogleTokens = googleTokens;
  }
}
