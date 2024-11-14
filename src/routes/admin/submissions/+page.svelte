<script>
	import { onMount } from 'svelte';
	import * as Icons from 'lucide-svelte';

	const API_BASE = 'http://localhost:8080/api';
	let submissions = [];
	let loading = true;
	let error = null;
	let selectedSubmission = null;

	async function fetchSubmissions() {
		try {
			const response = await fetch(`${API_BASE}/admin/submissions/all`, {
				headers: {
					Authorization: `Bearer ${localStorage.getItem('auth_token')}`
				}
			});
			const data = await response.json();
			submissions = data.data.submissions;
		} catch (err) {
			error = 'Failed to load submissions';
			console.error('Error:', err);
		} finally {
			loading = false;
		}
	}

	async function viewSubmission(id) {
		try {
			const response = await fetch(`${API_BASE}/admin/submissions/${id}`, {
				headers: {
					Authorization: `Bearer ${localStorage.getItem('auth_token')}`
				}
			});
			const data = await response.json();
			selectedSubmission = data.data;
		} catch (err) {
			console.error('Error fetching submission:', err);
		}
	}

	function formatDate(dateString) {
		return new Date(dateString).toLocaleString();
	}

	onMount(fetchSubmissions);
</script>

<div class="container mx-auto">
	<div class="mb-6 flex items-center justify-between">
		<h1 class="text-3xl font-bold">Submissions</h1>
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
							>Submission ID</th
						>
						<th
							class="px-6 py-3 text-left text-xs font-medium uppercase tracking-wider text-gray-500"
							>User ID</th
						>
						<th
							class="px-6 py-3 text-left text-xs font-medium uppercase tracking-wider text-gray-500"
							>Submitted</th
						>
						<th
							class="px-6 py-3 text-left text-xs font-medium uppercase tracking-wider text-gray-500"
							>Answers</th
						>
						<th
							class="px-6 py-3 text-left text-xs font-medium uppercase tracking-wider text-gray-500"
							>Actions</th
						>
					</tr>
				</thead>
				<tbody class="divide-y divide-gray-200 bg-white">
					{#each submissions as submission}
						<tr>
							<td class="whitespace-nowrap px-6 py-4">
								<div class="text-sm text-gray-900">{submission.id}</div>
							</td>
							<td class="whitespace-nowrap px-6 py-4">
								<div class="text-sm text-gray-900">{submission.user_id}</div>
							</td>
							<td class="whitespace-nowrap px-6 py-4">
								<div class="text-sm text-gray-500">{formatDate(submission.created_at)}</div>
							</td>
							<td class="whitespace-nowrap px-6 py-4">
								<div class="text-sm text-gray-900">{submission.answers.length} answers</div>
							</td>
							<td class="whitespace-nowrap px-6 py-4 text-sm text-gray-500">
								<button
									on:click={() => viewSubmission(submission.id)}
									class="text-blue-600 hover:text-blue-900"
								>
									View Details
								</button>
							</td>
						</tr>
					{/each}
				</tbody>
			</table>
		</div>
	{/if}

	{#if selectedSubmission}
		<div class="fixed inset-0 flex items-center justify-center bg-gray-500 bg-opacity-75">
			<div class="max-h-[90vh] w-full max-w-2xl overflow-y-auto rounded-lg bg-white p-6">
				<div class="mb-4 flex items-center justify-between">
					<h2 class="text-xl font-bold">Submission Details</h2>
					<button
						on:click={() => (selectedSubmission = null)}
						class="text-gray-500 hover:text-gray-700"
					>
						<Icons.X size={24} />
					</button>
				</div>

				<div class="space-y-4">
					<div class="grid grid-cols-2 gap-4 text-sm">
						<div>
							<span class="font-medium">Submission ID:</span>
							<span class="ml-2">{selectedSubmission.id}</span>
						</div>
						<div>
							<span class="font-medium">User ID:</span>
							<span class="ml-2">{selectedSubmission.user_id}</span>
						</div>
						<div>
							<span class="font-medium">Submitted:</span>
							<span class="ml-2">{formatDate(selectedSubmission.created_at)}</span>
						</div>
					</div>

					<div class="mt-4">
						<h3 class="mb-2 font-medium">Answers</h3>
						<div class="space-y-2">
							{#each selectedSubmission.answers as answer}
								<div class="border-b pb-2">
									<div class="font-medium">Question ID: {answer.id}</div>
									<div class="text-gray-600">Answer: {answer.question}</div>
								</div>
							{/each}
						</div>
					</div>
				</div>

				<div class="mt-6 flex justify-end">
					<button
						on:click={() => (selectedSubmission = null)}
						class="rounded bg-gray-200 px-4 py-2 text-gray-700 hover:bg-gray-300"
					>
						Close
					</button>
				</div>
			</div>
		</div>
	{/if}
</div>
