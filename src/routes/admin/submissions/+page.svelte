<script>
	import { onMount } from 'svelte';
	import * as Icons from 'lucide-svelte';

	const API_BASE = 'http://localhost:8080/api';
	let submissions = [];
	let questions = [];
	let loading = true;
	let error = null;
	let selectedSubmission = null;

	// Analytics data
	let questionStats = new Map();

	async function fetchData() {
		try {
			const [submissionsRes, questionsRes] = await Promise.all([
				fetch(`${API_BASE}/admin/submissions/all`, {
					headers: {
						Authorization: `Bearer ${localStorage.getItem('auth_token')}`
					}
				}),
				fetch(`${API_BASE}/questions`, {
					headers: {
						Authorization: `Bearer ${localStorage.getItem('auth_token')}`
					}
				})
			]);

			const submissionsData = await submissionsRes.json();
			const questionsData = await questionsRes.json();

			submissions = submissionsData.data.submissions.filter(s => 
				s.answers.some(a => a.question !== "")  // Filter out empty submissions
			);
			questions = questionsData.data.questions;

			calculateStats();
		} catch (err) {
			error = 'Failed to load data';
			console.error('Error:', err);
		} finally {
			loading = false;
		}
	}

	function calculateStats() {
		// Initialize statistics for each question
		questions.forEach(question => {
			questionStats.set(question.id, {
				question: question.question,
				totalResponses: 0,
				averageScore: 0,
				distribution: new Array(question.max - question.min + 1).fill(0),
				min: question.min,
				max: question.max,
				scores: []
			});
		});

		// Process all submissions
		submissions.forEach(submission => {
			submission.answers.forEach(answer => {
				const stats = questionStats.get(answer.id);
				if (stats && answer.question !== "") {
					const value = parseInt(answer.question);
					if (!isNaN(value)) {
						stats.scores.push(value);
						stats.totalResponses++;
						// Update distribution
						const index = value - stats.min;
						stats.distribution[index]++;
					}
				}
			});
		});

		// Calculate statistics
		questionStats.forEach(stats => {
			if (stats.totalResponses > 0) {
				stats.averageScore = stats.scores.reduce((a, b) => a + b, 0) / stats.totalResponses;
				stats.distribution = stats.distribution.map(count => ({
					count,
					percentage: (count / stats.totalResponses * 100).toFixed(1)
				}));
			}
		});

		questionStats = questionStats; // Trigger reactivity
	}

	function getMaxResponseCount(stats) {
		return Math.max(...stats.distribution.map(d => d.count));
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

	onMount(fetchData);
</script>

<div class="container mx-auto px-4 py-6">
	<div class="mb-6 flex items-center justify-between">
		<h1 class="text-3xl font-bold">Submissions Analysis</h1>
		<div class="text-sm text-gray-600">
			Total Submissions: {submissions.length}
		</div>
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
		<!-- Analytics Dashboard -->
		<div class="mb-8 grid grid-cols-1 gap-6 lg:grid-cols-2">
			{#each Array.from(questionStats.values()) as stats}
				<div class="rounded-lg bg-white p-6 shadow-lg">
					<h3 class="mb-4 text-lg font-semibold">{stats.question}</h3>
					<div class="mb-4 grid grid-cols-2 gap-4 text-sm">
						<div class="rounded-lg bg-blue-50 p-3">
							<p class="text-blue-600">Total Responses</p>
							<p class="text-2xl font-bold">{stats.totalResponses}</p>
						</div>
						<div class="rounded-lg bg-green-50 p-3">
							<p class="text-green-600">Average Score</p>
							<p class="text-2xl font-bold">{stats.averageScore.toFixed(1)}</p>
						</div>
					</div>
					
					<!-- Distribution Chart -->
					<div class="mt-6">
						<h4 class="mb-2 text-sm font-medium text-gray-600">Response Distribution</h4>
						<div class="space-y-2">
							{#each stats.distribution as dist, i}
								<div class="flex items-center">
									<div class="w-8 text-sm text-gray-600">{i + stats.min}</div>
									<div class="flex-1">
										<div class="h-6 w-full rounded-full bg-gray-100">
											<div
												class="h-6 rounded-full bg-blue-500 text-xs leading-6 text-white transition-all"
												style="width: {dist.percentage}%"
											>
												{#if dist.count > 0}
													<span class="ml-2">{dist.count} ({dist.percentage}%)</span>
												{/if}
											</div>
										</div>
									</div>
								</div>
							{/each}
						</div>
					</div>
				</div>
			{/each}
		</div>

		<!-- Submissions Table -->
		<div class="overflow-hidden rounded-lg bg-white shadow">
			<div class="p-6">
				<h2 class="mb-4 text-xl font-bold">Recent Submissions</h2>
				<table class="min-w-full">
					<thead class="bg-gray-50">
						<tr>
							<th class="px-6 py-3 text-left text-xs font-medium uppercase tracking-wider text-gray-500"
								>Submission ID</th>
							<th class="px-6 py-3 text-left text-xs font-medium uppercase tracking-wider text-gray-500"
								>User ID</th>
							<th class="px-6 py-3 text-left text-xs font-medium uppercase tracking-wider text-gray-500"
								>Submitted</th>
							<th class="px-6 py-3 text-left text-xs font-medium uppercase tracking-wider text-gray-500"
								>Answers</th>
							<th class="px-6 py-3 text-left text-xs font-medium uppercase tracking-wider text-gray-500"
								>Actions</th>
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
									<div class="text-sm text-gray-900">
										{submission.answers.filter(a => a.question !== "").length} answers
									</div>
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
							{#each selectedSubmission.answers.filter(a => a.question !== "") as answer}
								<div class="rounded-lg bg-gray-50 p-4">
									<div class="font-medium">{questions.find(q => q.id === answer.id)?.question || 'Unknown Question'}</div>
									<div class="text-gray-600">Score: {answer.question}</div>
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