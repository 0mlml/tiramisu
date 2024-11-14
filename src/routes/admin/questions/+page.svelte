<script>
	import { onMount } from 'svelte';
	import * as Icons from 'lucide-svelte';

	const API_BASE = 'http://localhost:8080/api';
	let questions = [];
	let loading = true;
	let error = null;
	let showAddModal = false;
	let newQuestion = {
		question: '',
		type: 'scale',
		min: 1,
		max: 5
	};

	async function fetchQuestions() {
		try {
			const response = await fetch(`${API_BASE}/questions`, {
				headers: {
					Authorization: `Bearer ${localStorage.getItem('auth_token')}`
				}
			});
			const data = await response.json();
			questions = data.data.questions;
		} catch (err) {
			error = 'Failed to load questions';
			console.error('Error:', err);
		} finally {
			loading = false;
		}
	}

	async function handleSubmit() {
		try {
			const response = await fetch(`${API_BASE}/admin/questions`, {
				method: 'POST',
				headers: {
					Authorization: `Bearer ${localStorage.getItem('auth_token')}`,
					'Content-Type': 'application/json'
				},
				body: JSON.stringify(newQuestion)
			});

			if (response.ok) {
				showAddModal = false;
				newQuestion = {
					question: '',
					type: 'scale',
					min: 1,
					max: 5
				};
				await fetchQuestions();
			}
		} catch (err) {
			console.error('Error creating question:', err);
		}
	}

	async function deleteQuestion(id) {
		if (confirm('Are you sure you want to delete this question?')) {
			try {
				const response = await fetch(`${API_BASE}/admin/questions/${id}`, {
					method: 'DELETE',
					headers: {
						Authorization: `Bearer ${localStorage.getItem('auth_token')}`
					}
				});

				if (response.ok) {
					await fetchQuestions();
				}
			} catch (err) {
				console.error('Error deleting question:', err);
			}
		}
	}

	onMount(fetchQuestions);
</script>

<div class="container mx-auto">
	<div class="mb-6 flex items-center justify-between">
		<h1 class="text-3xl font-bold">Questions</h1>
		<button
			on:click={() => (showAddModal = true)}
			class="rounded bg-blue-500 px-4 py-2 text-white hover:bg-blue-600"
		>
			<div class="flex items-center">
				<Icons.Plus class="mr-2" size={20} />
				Add Question
			</div>
		</button>
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
							>Question</th
						>
						<th
							class="px-6 py-3 text-left text-xs font-medium uppercase tracking-wider text-gray-500"
							>Type</th
						>
						<th
							class="px-6 py-3 text-left text-xs font-medium uppercase tracking-wider text-gray-500"
							>Scale</th
						>
						<th
							class="px-6 py-3 text-left text-xs font-medium uppercase tracking-wider text-gray-500"
							>Actions</th
						>
					</tr>
				</thead>
				<tbody class="divide-y divide-gray-200 bg-white">
					{#each questions as question}
						<tr>
							<td class="px-6 py-4">
								<div class="text-sm text-gray-900">{question.question}</div>
							</td>
							<td class="whitespace-nowrap px-6 py-4">
								<span
									class="inline-flex rounded-full bg-blue-100 px-2 text-xs font-semibold leading-5 text-blue-800"
								>
									{question.type}
								</span>
							</td>
							<td class="whitespace-nowrap px-6 py-4 text-sm text-gray-500">
								{question.min} - {question.max}
							</td>
							<td class="whitespace-nowrap px-6 py-4 text-right text-sm font-medium">
								<button
									on:click={() => deleteQuestion(question.id)}
									class="text-red-600 hover:text-red-900"
								>
									Delete
								</button>
							</td>
						</tr>
					{/each}
				</tbody>
			</table>
		</div>
	{/if}

	{#if showAddModal}
		<div class="fixed inset-0 flex items-center justify-center bg-gray-500 bg-opacity-75">
			<div class="w-full max-w-lg rounded-lg bg-white p-6">
				<h2 class="mb-4 text-xl font-bold">Add New Question</h2>
				<form on:submit|preventDefault={handleSubmit}>
					<div class="mb-4">
						<label class="block text-sm font-medium text-gray-700">Question</label>
						<input
							type="text"
							bind:value={newQuestion.question}
							class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
							required
						/>
					</div>
					<div class="mb-4 grid grid-cols-2 gap-4">
						<div>
							<label class="block text-sm font-medium text-gray-700">Min Value</label>
							<input
								type="number"
								bind:value={newQuestion.min}
								class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
								required
							/>
						</div>
						<div>
							<label class="block text-sm font-medium text-gray-700">Max Value</label>
							<input
								type="number"
								bind:value={newQuestion.max}
								class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
								required
							/>
						</div>
					</div>
					<div class="flex justify-end space-x-3">
						<button
							type="button"
							on:click={() => (showAddModal = false)}
							class="rounded bg-gray-200 px-4 py-2 text-gray-700 hover:bg-gray-300"
						>
							Cancel
						</button>
						<button
							type="submit"
							class="rounded bg-blue-500 px-4 py-2 text-white hover:bg-blue-600"
						>
							Add Question
						</button>
					</div>
				</form>
			</div>
		</div>
	{/if}
</div>
