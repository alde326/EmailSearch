import { Root } from "../types/emailTypes";

import {
  SERVER,
  GET_SINGLE_MAIL_ENDPOINT,
  METHOD_GET,
  AUTHORIZATION_LEVEL,
  USER,
  PASSWORD,
} from "../constants/constants";

const getSingleEmail = async (emailID: string) => {
  const url = new URL(`${SERVER}${GET_SINGLE_MAIL_ENDPOINT}`);
  url.searchParams.append("id", emailID);
  try {
    const response = await fetch(url, {
      method: METHOD_GET,
      headers: {
        Authorization: AUTHORIZATION_LEVEL + btoa(`${USER}:${PASSWORD}`),
      },
    });
    if (!response.ok) {
      throw new Error(
        `Error in your request to Get Emails, status: ${response?.status}`
      );
    }
    const data: Root = await response.json();
    if (data.hits.hits.length === 1) return data?.hits.hits[0]._source;
    throw new Error(
        `There's an error with the amount of emails, status: ${response?.status}`
      );
  } catch (error) {
    console.error("There was a problem with the fetch operation:", error);
  }
};

export default getSingleEmail;
