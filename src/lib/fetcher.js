const BASE_URL = "http://localhost:8080/api"

export const loginUser = async function (email, pass) {
    try {
        const response = await fetch(BASE_URL + "/login", {
            method: "POST",
            body: JSON.stringify({
                email: email,
                password: pass
            }),
            headers: {
                "Content-Type": "application/json; charset=UTF-8"
            }
        });

        if (!response.ok) {
            throw new Error(`Error: ${response.status} ${response.statusText}`);
        }

        return await response.json();

    } catch (error) {
        console.error("Login failed:", error);
        throw error;
    }
};

export const getProfile = async function (token) {
    try {
        const response = await fetch(BASE_URL + "/profile", {
            method: "GET",
            headers: {
                "Authorization": `Bearer ${token}`,
                "Content-Type": "application/json; charset=UTF-8"
            }
        });

        if (!response.ok) {
            throw new Error(`Error: ${response.status} ${response.statusText}`);
        }

        return await response.json();

    } catch (error) {
        console.error("Getting profile failed:", error);
        throw error;
    }
};