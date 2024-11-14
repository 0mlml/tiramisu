import { redirect } from '@sveltejs/kit';

export async function load({ cookies, fetch }) {
    const token = cookies.get('auth_token');

    if (!token) {
        throw redirect(303, '/sign-in');
    }

    try {
        const response = await fetch('http://localhost:8080/api/profile', {
            headers: {
                'Authorization': `Bearer ${token}`
            }
        });

        const userData = await response.json();

        if (!userData.data.is_admin) {
            throw redirect(303, '/');
        }

        return {
            user: userData.data
        };
    } catch (error) {
        throw redirect(303, '/sign-in');
    }
}