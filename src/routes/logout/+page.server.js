import { redirect } from '@sveltejs/kit';
import { expoIn } from 'svelte/easing';

export const actions = {
    default: async ({ cookies, fetch }) => {
        cookies.delete('auth_token', {
            path: '/'
        });
        throw redirect(303, '/sign-in');
    }
};