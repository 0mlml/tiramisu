import { redirect } from "@sveltejs/kit";
import * as fetcher from "$lib/fetcher.js";

export const load = async (event) => {
  const token = event.cookies.get("auth_token");

  if (!token) {
    throw redirect(301, "/sign-in");
  }

  const data = await fetcher.getProfile(token)['data']
  return data;
};