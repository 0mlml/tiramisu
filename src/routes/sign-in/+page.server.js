import * as fetcher from "$lib/fetcher";
import { error, redirect } from "@sveltejs/kit";

export const load = async (event) => {
  const token = event.cookies.get("auth_token");

  if (token && token !== "") {
    throw redirect(301, "/home");
  }
};

export const actions = {
  default: async (event) => {
    const formData = await event.request.formData();
    const email = formData.get("email");
    const password = formData.get("password");

    if (!email || !password) {
      throw error(400, "must provide an email and password");
    }

    try{
        const token = (await fetcher.loginUser(email, password))["data"]["token"];
        event.cookies.set("auth_token", token ?? "", {
            path: "/",
        });    
    }catch (err) {
        console.error("Error fetching token", err);
    }

    throw redirect(301, "/home");
  },
};