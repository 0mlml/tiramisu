import { serverAuth } from "$lib/server/api-client";

export const logout = async (event) => {
    serverAuth.logout(event);
  }