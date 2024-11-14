const BASE_URL = "http://localhost:8080/api";
import { redirect } from "@sveltejs/kit";

export class ApiError extends Error {
    constructor(status, statusText, data) {
        super(`${status} ${statusText}`);
        this.status = status;
        this.statusText = statusText;
        this.data = data;
    }
}

async function serverFetch(event, endpoint, options = {}) {
    const response = await event.fetch(`${BASE_URL}${endpoint}`, {
        ...options,
        headers: {
            'Content-Type': 'application/json',
            ...options.headers,
        },
    });



    const data = await response.json();

    if (!response.ok) {
        throw new ApiError(response.status, response.statusText, data);
    }

    return data;
}

export const serverAuth = {
    login: async (event, userData) => {
        return serverFetch(event, '/login', {
            method: 'POST',
            body: JSON.stringify(userData),
        });
    },

    validateToken: async (event, token) => {
        return serverFetch(event, '/profile', {
            headers: {
                Authorization: `Bearer ${token}`,
            },
        });
    },

    register: async (event, userData) => {
        return serverFetch(event, '/register', {
            method: 'POST',
            body: JSON.stringify(userData),
        });
    },

    logout: async (event) => {
        localStorage.removeItem('auth_token');
        event.cookies.set('auth_token', '', {});
        throw redirect(303, '/sign-in');
    }
};