import { RouteLocationNormalizedLoaded } from "vue-router";
import { Source } from "../types/emailTypes";

export const getParsedEmailId = (route: RouteLocationNormalizedLoaded) => {
  let content: string | string[] = route.params.email_id;
  if (Array.isArray(content)) {
    content = content.join(" ");
  }
  return content;
};

export const getEmailFooterDetails = (
  parsedEmail: Source | null | undefined
) => [
  { detailName: "XBcc", detailValue: parsedEmail?.x_bcc },
  { detailName: "XOrigin", detailValue: parsedEmail?.x_origin },
  { detailName: "XFilename", detailValue: parsedEmail?.x_file_name },
  { detailName: "CFolder", detailValue: parsedEmail?.c_folder },
  { detailName: "MimeVersion", detailValue: parsedEmail?.mime_version },
  { detailName: "ContentType", detailValue: parsedEmail?.content_type },
  {
    detailName: "contentenconding",
    detailValue: parsedEmail?.content_transfer_encoding,
  },
];

export const getEmailLastHeaderDetails = (
  parsedEmail: Source | null | undefined
) => [
  {
    detailName: "From",
    detailFirstValue: parsedEmail?.x_from,
    detailSecondValue: parsedEmail?.from,
  },
  {
    detailName: "To",
    detailFirstValue: parsedEmail?.x_to,
    detailSecondValue: parsedEmail?.to,
  },
];
