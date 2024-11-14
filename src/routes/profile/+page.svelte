<script>
	import { onMount } from 'svelte';
	import * as Icons from 'lucide-svelte';
	import { goto } from '$app/navigation';

	const API_BASE = 'http://localhost:8080/api';
	let profile = null;
	let loading = true;
	let error = null;
	let success = null;
	let isEditing = false;
	let editForm = {
		name: '',
		picture: ''
	};

	async function fetchProfile() {
		try {
			const response = await fetch(`${API_BASE}/profile`, {
				headers: {
					Authorization: `Bearer ${localStorage.getItem('auth_token')}`
				}
			});

			if (!response.ok) {
				if (response.status === 401) {
					goto('/sign-in');
					return;
				}
				throw new Error('Failed to fetch profile');
			}

			const data = await response.json();
			profile = data.data;
			editForm.name = profile.name;
			editForm.picture = profile.picture || '';
		} catch (err) {
			error = 'Failed to load profile';
			console.error('Error:', err);
		} finally {
			loading = false;
		}
	}

	async function handleSubmit() {
		try {
			const response = await fetch(`${API_BASE}/profile`, {
				method: 'PUT',
				headers: {
					Authorization: `Bearer ${localStorage.getItem('auth_token')}`,
					'Content-Type': 'application/json'
				},
				body: JSON.stringify(editForm)
			});

			if (!response.ok) {
				throw new Error('Failed to update profile');
			}

			success = 'Profile updated successfully';
			await fetchProfile();
			isEditing = false;

			// Clear success message after 3 seconds
			setTimeout(() => {
				success = null;
			}, 3000);
		} catch (err) {
			error = 'Failed to update profile';
			console.error('Error:', err);
		}
	}

	onMount(fetchProfile);
</script>

<div class="container mx-auto px-4 py-8">
	<div class="mx-auto max-w-2xl">
		<h1 class="mb-8 text-3xl font-bold">Profile</h1>

		{#if loading}
			<div class="flex h-64 items-center justify-center">
				<p class="text-lg">Loading...</p>
			</div>
		{:else if error}
			<div class="mb-4 rounded border border-red-400 bg-red-100 px-4 py-3 text-red-700">
				{error}
			</div>
		{:else if success}
			<div class="mb-4 rounded border border-green-400 bg-green-100 px-4 py-3 text-green-700">
				{success}
			</div>
		{:else}
			<div class="overflow-hidden rounded-lg bg-white shadow">
				{#if !isEditing}
					<!-- View Mode -->
					<div class="p-6">
						<div class="flex items-start justify-between">
							<div class="flex items-center space-x-4">
								{#if profile.picture}
									<img src={profile.picture} alt={profile.name} class="h-20 w-20 rounded-full" />
								{:else}
									<div class="flex h-20 w-20 items-center justify-center rounded-full bg-gray-200">
										<span class="text-2xl font-medium text-gray-600">
											{profile.name[0].toUpperCase()}
										</span>
									</div>
								{/if}
								<div>
									<h2 class="text-2xl font-bold">{profile.name}</h2>
									<p class="text-gray-600">
										Member since {profile.created}
									</p>
									{#if profile.is_admin}
										<span
											class="mt-2 inline-flex items-center rounded-full bg-blue-100 px-2.5 py-0.5 text-xs font-medium text-blue-800"
										>
											Administrator
										</span>
									{/if}
								</div>
							</div>
							<button
								on:click={() => (isEditing = true)}
								class="flex items-center text-blue-600 hover:text-blue-800"
							>
								<Icons.Edit2 size={20} class="mr-1" />
								Edit Profile
							</button>
						</div>
					</div>
				{:else}
					<!-- Edit Mode -->
					<form on:submit|preventDefault={handleSubmit} class="space-y-4 p-6">
						<div>
							<label for="name" class="block text-sm font-medium text-gray-700">Name</label>
							<input
								type="text"
								id="name"
								bind:value={editForm.name}
								class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
								required
							/>
						</div>

						<div>
							<label for="picture" class="block text-sm font-medium text-gray-700"
								>Profile Picture URL</label
							>
							<input
								type="url"
								id="picture"
								bind:value={editForm.picture}
								class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
								placeholder="https://example.com/picture.jpg"
							/>
						</div>

						<div class="flex justify-end space-x-3">
							<button
								type="button"
								on:click={() => (isEditing = false)}
								class="rounded-md border border-gray-300 bg-white px-4 py-2 text-sm font-medium text-gray-700 shadow-sm hover:bg-gray-50"
							>
								Cancel
							</button>
							<button
								type="submit"
								class="rounded-md border border-transparent bg-blue-600 px-4 py-2 text-sm font-medium text-white shadow-sm hover:bg-blue-700"
							>
								Save Changes
							</button>
						</div>
					</form>
				{/if}
			</div>
		{/if}
	</div>
</div>
