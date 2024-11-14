import * as fetcher from "$lib/fetcher";
import { error, redirect } from "@sveltejs/kit";

export const load = async (event) => {
  const token = event.cookies.get("auth_token");

  if (token && token !== "") {
    throw redirect(301, "/");
  }
};

export const actions = {
  default: async (event) => {
    const formData = await event.request.formData();
    const email = formData.get("email");
    const password = formData.get("password");

    if (!email || !password) {
      throw error(400, "Must provide an email and password");
    }

    const response = await fetcher.loginUser(email, password).catch((error) => {
      throw error(400, "Invalid email or password");
    });
    const token = response.data.token;

    event.cookies.set("auth_token", token ?? "", {
      path: "/",
      httpOnly: true,
      sameSite: "lax",
      secure: process.env.NODE_ENV === "production"
    });

    return { success: true, token };
  },
};