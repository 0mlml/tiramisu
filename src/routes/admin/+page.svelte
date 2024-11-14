<script>
	import { onMount } from 'svelte';

	let submissions = [];
	let questions = [];
	let selectedQuestion = null;
	let submissionStats = [];
	let loading = true;
	const API_BASE = 'http://localhost:8080/api';

	async function fetchData() {
		try {
			const [submissionsRes, questionsRes] = await Promise.all([
				fetch(API_BASE + '/admin/submissions/all', {
					headers: {
						Authorization: `Bearer ${localStorage.getItem('auth_token')}`
					}
				}),
				fetch(API_BASE + '/questions', {
					headers: {
						Authorization: `Bearer ${localStorage.getItem('auth_token')}`
					}
				})
			]);

			const submissionsData = await submissionsRes.json();
			const questionsData = await questionsRes.json();

			submissions = submissionsData.data.submissions || [];
			questions = questionsData.data.questions || [];

			processSubmissionStats();
		} catch (error) {
			console.error('Error fetching data:', error);
		} finally {
			loading = false;
		}
	}

	function processSubmissionStats() {
		if (!selectedQuestion) return;

		const stats = submissions.reduce((acc, submission) => {
			const answer = submission.answers.find((a) => a.id === selectedQuestion.id);
			if (answer) {
				const value = parseInt(answer.question);
				const date = new Date(submission.created_at).toLocaleDateString();

				const existingEntry = acc.find((entry) => entry.date === date);
				if (existingEntry) {
					existingEntry.values.push(value);
					existingEntry.average =
						existingEntry.values.reduce((a, b) => a + b, 0) / existingEntry.values.length;
				} else {
					acc.push({
						date,
						values: [value],
						average: value
					});
				}
			}
			return acc;
		}, []);

		submissionStats = stats.sort((a, b) => new Date(a.date) - new Date(b.date));
	}

	onMount(() => {
		fetchData();
	});

	$: if (selectedQuestion) {
		processSubmissionStats();
	}
</script>

<div class="container mx-auto">
	<h1 class="mb-8 text-3xl font-bold">Admin Dashboard</h1>

	{#if loading}
		<div class="flex h-64 items-center justify-center">
			<p class="text-lg">Loading...</p>
		</div>
	{:else}
		<div class="mb-8 grid grid-cols-1 gap-6 md:grid-cols-2">
			<div class="rounded-lg bg-white p-6 shadow">
				<h2 class="mb-4 text-xl font-semibold">Submission Statistics</h2>
				<p class="text-4xl font-bold">{submissions.length}</p>
				<p class="text-gray-600">Total Submissions</p>
			</div>

			<div class="rounded-lg bg-white p-6 shadow">
				<h2 class="mb-4 text-xl font-semibold">Question Statistics</h2>
				<p class="text-4xl font-bold">{questions.length}</p>
				<p class="text-gray-600">Total Questions</p>
			</div>
		</div>

		<div class="rounded-lg bg-white p-6 shadow">
			<h2 class="mb-4 text-xl font-semibold">Recent Submissions</h2>
			<div class="overflow-x-auto">
				<table class="w-full">
					<thead>
						<tr>
							<th class="p-2 text-left">Date</th>
							<th class="p-2 text-left">User ID</th>
							<th class="p-2 text-left">Answers</th>
						</tr>
					</thead>
					<tbody>
						{#each submissions.slice(0, 10) as submission}
							<tr class="border-t">
								<td class="p-2">
									{new Date(submission.created_at).toLocaleDateString()}
								</td>
								<td class="p-2">{submission.user_id}</td>
								<td class="p-2">{submission.answers.length} responses</td>
							</tr>
						{/each}
					</tbody>
				</table>
			</div>
		</div>
	{/if}
</div>
