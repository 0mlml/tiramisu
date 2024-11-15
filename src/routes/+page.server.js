import { redirect } from "@sveltejs/kit";
import * as fetcher from "$lib/fetcher.js";
import { linkUtil } from "$lib/linkUtil.js";

export const load = async (event) => {
  const token = event.cookies.get("auth_token") || null;

  if (token === null){
    redirect(302, '/sign-in');
    return
  }

  let profile = null;
    
  if (token !== null){
    profile = (await fetcher.getProfile(token))["data"];
  }
  
  if (profile !== null){
    linkUtil["links"][0]["linkText"] = profile["name"];
    if (profile["picture"] !== ""){
      linkUtil["links"][0]["picture"] = profile["picture"];
    }
    } else {
    linkUtil["links"][0]["linkText"] = "Sign in";
    linkUtil["links"][0]["url"] = "/sign-in" 
  }

  return { linkUtil };
};