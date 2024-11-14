import { redirect } from '@sveltejs/kit';

export async function load({ cookies, fetch }) {
    const token = cookies.get('auth_token');
    
    if (!token) {
        throw redirect(303, '/sign-in');
    }

    return {
        token
    };
}