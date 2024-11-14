<script>
	import { onMount } from 'svelte';

	onMount(() => {
		localStorage.removeItem('auth_token');
	});
</script>

<div class="flex min-h-screen items-center justify-center bg-gray-100">
	<div class="text-center">
		<h1 class="mb-4 text-2xl font-semibold">Logging out...</h1>
		<p class="text-gray-600">Please wait while we sign you out.</p>
	</div>
</div>
