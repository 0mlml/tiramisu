<script>
    import { goto } from '$app/navigation';
</script>

<main class="w-full min-h-screen flex items-center justify-center bg-neutral-100">
    <div class="bg-white rounded-lg shadow-md p-8 text-center max-w-md">
        <div class="mb-6">
            <svg class="w-16 h-16 text-emerald-500 mx-auto" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path>
            </svg>
        </div>
        
        <h1 class="text-2xl font-medium mb-4">Thank You!</h1>
        <p class="text-neutral-600 mb-8">Your responses have been successfully submitted.</p>
        
        <form class="entry-form" action="/logout" method="POST">
            <button
                type="submit"
                class="px-6 py-2 bg-emerald-500 text-white rounded-full hover:bg-emerald-600 transition-colors"
                >
                Logout
            </button>
        </form>
    </div>
</main>