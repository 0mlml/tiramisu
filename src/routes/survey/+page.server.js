import { serverAuth } from '$lib/server/api-client';
import { error, redirect } from '@sveltejs/kit';

export const load = async ({ cookies }) => {
    const token = cookies.get('auth_token');

    if (!token) {
        throw redirect(303, '/sign-in');
    }

    try {
        const response = await fetch('http://localhost:8080/api/questions', {
            headers: {
                'Authorization': `Bearer ${token}`
            }
        });

        if (!response.ok) {
            throw error(response.status, 'Failed to fetch questions');
        }

        const data = await response.json();
        return {
            questions: data.data.questions
        };
    } catch (err) {
        console.error('Error fetching questions:', err);
        throw error(500, 'Failed to load questionnaire');
    }
};

export const actions = {
    submit: async ({ request, cookies }) => {
        const token = cookies.get('auth_token');
        const formData = await request.formData();
        
        // Convert form data to the expected format
        const answers = [];
        for (const [key, value] of formData.entries()) {
            if (key.startsWith('question_')) {
                const id = key.replace('question_', '');
                answers.push({
                    id,
                    value: value
                });
            }
        }

        try {
            const response = await fetch('http://localhost:8080/api/submit', {
                method: 'POST',
                headers: {
                    'Authorization': `Bearer ${token}`,
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({ answers })
            });

            if (!response.ok) {
                throw new Error('Failed to submit questionnaire');
            }

            const result = await response.json();
            return { success: true, submissionId: result.data.id };
        } catch (err) {
            console.error('Error submitting questionnaire:', err);
            return { success: false, error: err.message };
        }
    }
};