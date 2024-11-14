import { serverAuth, ApiError } from '$lib/server/api-client';
import { error, redirect } from '@sveltejs/kit';
import { dev } from '$app/environment';

export const load = async ({ cookies }) => {
    const token = cookies.get('auth_token');

    if (token) {
        throw redirect(303, '/');
    }
};

export const actions = {
    default: async (event) => {
        const formData = await event.request.formData();
        const email = formData.get('email');
        const password = formData.get('password');
        const confirmPassword = formData.get('confirmPassword');
        const name = formData.get('name');

        if (!email || !password || !confirmPassword || !name) {
            throw error(400, {
                message: 'All fields are required'
            });
        }

        if (password !== confirmPassword) {
            throw error(400, {
                message: 'Passwords do not match'
            });
        }

        if (password.length < 8) {
            throw error(400, {
                message: 'Password must be at least 8 characters long'
            });
        }

        try {
            const response = await serverAuth.register(event, { email, password, name });
            const token = response.data.token;

            event.cookies.set('auth_token', token, {
                path: '/',
                httpOnly: true,
                secure: !dev,
                sameSite: 'lax',
                maxAge: 60 * 60 * 24 * 7,
            });

            redirect(303, '/');

            return { success: true };
        } catch (err) {
            console.error('Registration error:', err);

            if (err instanceof ApiError && err.status === 400) {
                throw error(400, {
                    message: err.data.data || 'Email already exists'
                });
            }

            throw error(500, {
                message: 'An unexpected error occurred'
            });
        }
    }
};