<script>
	import { onMount } from 'svelte';
	import * as Icons from 'lucide-svelte';

	const API_BASE = 'http://localhost:8080/api';
	let users = [];
	let loading = true;
	let error = null;

	async function fetchUsers() {
		try {
			const response = await fetch(`${API_BASE}/admin/users`, {
				headers: {
					Authorization: `Bearer ${localStorage.getItem('auth_token')}`
				}
			});
			const data = await response.json();
			users = [...data.data.users];
		} catch (err) {
			error = 'Failed to load users';
			console.error('Error:', err);
		} finally {
			loading = false;
		}
	}

	onMount(fetchUsers);
</script>

<div class="container mx-auto">
	<div class="mb-6 flex items-center justify-between">
		<h1 class="text-3xl font-bold">Users</h1>
	</div>

	{#if loading}
		<div class="flex h-64 items-center justify-center">
			<p class="text-lg">Loading...</p>
		</div>
	{:else if error}
		<div class="rounded border border-red-400 bg-red-100 px-4 py-3 text-red-700">
			{error}
		</div>
	{:else}
		<div class="overflow-hidden rounded-lg bg-white shadow">
			<table class="min-w-full">
				<thead class="bg-gray-50">
					<tr>
						<th
							class="px-6 py-3 text-left text-xs font-medium uppercase tracking-wider text-gray-500"
							>User</th
						>
						<th
							class="px-6 py-3 text-left text-xs font-medium uppercase tracking-wider text-gray-500"
							>Email</th
						>
						<th
							class="px-6 py-3 text-left text-xs font-medium uppercase tracking-wider text-gray-500"
							>Admin Status</th
						>
						<th
							class="px-6 py-3 text-left text-xs font-medium uppercase tracking-wider text-gray-500"
							>Created</th
						>
					</tr>
				</thead>
				<tbody class="divide-y divide-gray-200 bg-white">
					{#each users as user}
						<tr>
							<td class="whitespace-nowrap px-6 py-4">
								<div class="flex items-center">
									{#if user.picture}
										<img class="h-10 w-10 rounded-full" src={user.picture} alt="" />
									{:else}
										<div
											class="flex h-10 w-10 items-center justify-center rounded-full bg-gray-200"
										>
											<span class="text-sm font-medium text-gray-500">
												{user.name[0].toUpperCase()}
											</span>
										</div>
									{/if}
									<div class="ml-4">
										<div class="text-sm font-medium text-gray-900">{user.name}</div>
									</div>
								</div>
							</td>
							<td class="whitespace-nowrap px-6 py-4">
								<div class="text-sm text-gray-900">{user.email}</div>
							</td>
							<td class="whitespace-nowrap px-6 py-4">
								<button
									class="inline-flex items-center rounded-full px-3 py-1 text-sm font-medium {user.is_admin
										? 'bg-green-100 text-green-800'
										: 'bg-gray-100 text-gray-800'}"
								>
									{user.is_admin ? 'Admin' : 'User'}
								</button>
							</td>
							<td class="whitespace-nowrap px-6 py-4 text-sm text-gray-500">
								<!-- {new Date(user.created).toLocaleDateString()} -->
                                 {user.created}
							</td>
						</tr>
					{/each}
				</tbody>
			</table>
		</div>
	{/if}
</div>
