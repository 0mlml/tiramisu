<script>
	import { enhance } from '$app/forms';
	import { goto } from '$app/navigation';
	export let data;
	export let form;

	let loading = false;

	function handleSubmit() {
		loading = true;

		return async ({result}) => {
			loading = false;
      		localStorage.setItem('auth_token', result.data.token);
			window.location.reload();
		};
	}
</script>

<main class="flex h-screen w-full items-center justify-center bg-neutral-100">
	<form
		method="POST"
		use:enhance={handleSubmit}
		class="w-[400px] rounded-sm bg-white p-8 shadow-md"
	>
		<h2 class="mb-8 text-center text-2xl text-neutral-950">Sign In</h2>

		{#if form?.error}
			<div class="mb-4 rounded-lg border border-red-200 bg-red-50 p-3">
				<p class="text-center text-sm font-light text-red-600">
					{form.error.message || 'Invalid credentials'}
				</p>
			</div>
		{/if}

		<div class="mb-2">
			<label for="email" class="mb-1 ml-2 block font-light text-neutral-800"> Email </label>
			<input
				type="email"
				id="email"
				name="email"
				placeholder="Enter your email"
				class="w-full rounded-lg border border-neutral-300 p-2 transition-colors focus:border-emerald-400 focus:outline-none focus:ring-2 focus:ring-emerald-200"
				required
				disabled={loading}
			/>
		</div>

		<div class="mb-6">
			<label for="password" class="mb-1 ml-2 block font-light text-neutral-800"> Password </label>
			<input
				class="w-full rounded-lg border border-neutral-300 p-2 transition-colors focus:border-emerald-400 focus:outline-none focus:ring-2 focus:ring-emerald-200"
				type="password"
				id="password"
				name="password"
				placeholder="Enter your password"
				required
				disabled={loading}
			/>
		</div>

		<div class="flex w-full justify-center">
			<button
				type="submit"
				class="submit rounded-full px-6 py-2 text-lg text-white transition-all duration-200 hover:bg-emerald-500 focus:outline-none focus:ring-2 focus:ring-emerald-400 focus:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50"
				disabled={loading}
			>
				{#if loading}
					<span class="flex items-center">
						<svg
							class="-ml-1 mr-2 h-4 w-4 animate-spin text-white"
							xmlns="http://www.w3.org/2000/svg"
							fill="none"
							viewBox="0 0 24 24"
						>
							<circle
								class="opacity-25"
								cx="12"
								cy="12"
								r="10"
								stroke="currentColor"
								stroke-width="4"
							></circle>
							<path
								class="opacity-75"
								fill="currentColor"
								d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"
							></path>
						</svg>
						Signing in...
					</span>
				{:else}
					Sign In
				{/if}
			</button>
		</div>

		<div class="mt-8 flex w-full justify-center">
			<a
				href="/sign-up"
				class="font-light text-neutral-600 transition-colors hover:text-emerald-600"
			>
				Create Account
			</a>
		</div>
	</form>
</main>

<style>
	.submit {
		background-color: #63c19d;
	}

	/* Optional: Add styles for the loading spinner */
	@keyframes spin {
		to {
			transform: rotate(360deg);
		}
	}

	.animate-spin {
		animation: spin 1s linear infinite;
	}
</style>
