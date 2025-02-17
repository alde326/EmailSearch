import { Root } from "../types/emailTypes";

import {
  AUTHORIZATION_LEVEL,
  GET_SEARCHED_MAILS_ENDPOINT,
  METHOD_GET,
  SERVER,
  USER,
  PASSWORD,
  MATCH_PHRASE,
  MATCH_ALL,
} from "../constants/constants";

const search = async (
  searchTerm: string,
  page: string,
  orderDate: string,
  isSearching: boolean = false
) => {
  const url = new URL(`${SERVER}${GET_SEARCHED_MAILS_ENDPOINT}`);
  isSearching
    ? url.searchParams.append("type", MATCH_PHRASE)
    : url.searchParams.append("type", MATCH_ALL);
  url.searchParams.append("term", searchTerm);
  url.searchParams.append("page", page);
  url.searchParams.append("order", orderDate);

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
    if (data) return data?.hits;
  } catch (error) {
    console.error("There was a problem with the fetch operation:", error);
  }
};

export default search;
