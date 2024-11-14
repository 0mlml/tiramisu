import * as fetcher from "$lib/fetcher.js";
import { linkUtil } from "$lib/linkUtil.js";

export const load = async (event) => {
    const token = event.cookies.get("auth_token") || null;

    let profile = null;
  
    if (token !== null){
        profile = (await fetcher.getProfile(token))["data"];
    }

    if (profile !== null){
		linkUtil["links"][0]["linkText"] = profile["name"];
		linkUtil["links"][0]["url"] = "/profile" 
		if (profile["picture"] !== ""){
			linkUtil["links"][0]["picture"] = profile["picture"];
		}
	} else {
		linkUtil["links"][0]["linkText"] = "Sign in";
		linkUtil["links"][0]["url"] = "/sign-in" 
	}
    return { linkUtil };
};
