import * as fetcher from "$lib/fetcher.js";

export const load = async (event) => {
    const token = event.cookies.get("auth_token") || null;

    let profile = null;
  
    if (token !== null){
        profile = (await fetcher.getProfile(token))['data'];
    }

    return { profile };
};