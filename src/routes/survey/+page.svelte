<script>
	import { enhance } from '$app/forms';
	import { goto } from '$app/navigation';

	export let data;
	export let form;

	let loading = false;
	let answers = new Map();

	$: questions = data.questions || [];
	$: isSubmitEnabled = answers.size === questions.length;

	function handleAnswerChange(questionId, value) {
		answers.set(questionId, value);
		answers = answers;
	}

	function handleSubmit() {
		loading = true;
		
		const formAnswers = Array.from(answers.entries()).map(([id, value]) => ({
			id,
			question: questions.find(q => q.id === id)?.question || ''
		}));

		const hiddenInput = document.createElement('input');
		hiddenInput.type = 'hidden';
		hiddenInput.name = 'answers';
		hiddenInput.value = JSON.stringify(formAnswers);
		event.target.appendChild(hiddenInput);

		return async ({ result }) => {
			loading = false;
			if (result.type === 'success') {
				goto('/survey/complete');
			}
		};
	}

	function getScaleLabels(min, max) {
		return Array.from({ length: max - min + 1 }, (_, i) => min + i);
	}
</script>

<main class="min-h-screen w-full bg-neutral-100 px-4 py-8">
	<div class="mx-auto max-w-2xl">
		<div class="mb-8">
			<h1 class="text-2xl font-bold text-neutral-800">Questionnaire</h1>
			<p class="mt-2 text-neutral-600">Please answer all questions below</p>
		</div>

		<div class="rounded-lg bg-white p-8 shadow-md">
			{#if form?.error}
				<div class="mb-6 rounded-lg border border-red-200 bg-red-50 p-4">
					<p class="text-center text-red-600">{form.error}</p>
				</div>
			{/if}

			<form method="POST" action="?/submit" use:enhance={handleSubmit} class="space-y-12">
				{#if questions.length > 0}
					{#each questions as question, index}
						<div class="border-b border-neutral-200 pb-8 last:border-0">
							<div class="mb-6">
								<h2 class="text-xl font-medium text-neutral-800">
									{index + 1}. {question.question}
								</h2>
							</div>

							{#if question.type === 'scale'}
								<div class="flex flex-col items-center">
									<div class="mb-2 flex w-full max-w-md justify-between">
										{#each getScaleLabels(question.min, question.max) as value}
											<label class="group flex cursor-pointer flex-col items-center">
												<input
													type="radio"
													name="question_{question.id}"
													{value}
													checked={answers.get(question.id) === value}
													class="peer sr-only"
													on:change={() => handleAnswerChange(question.id, value)}
												/>
												<span
													class="flex h-10 w-10 items-center justify-center rounded-full border-2 border-neutral-300 text-neutral-600 transition-all duration-200 group-hover:border-emerald-400 peer-checked:border-emerald-500 peer-checked:bg-emerald-500 peer-checked:text-white"
												>
													{value}
												</span>
											</label>
										{/each}
									</div>
									<div class="flex w-full max-w-md justify-between text-sm text-neutral-500">
										<span>Not at all</span>
										<span>Very much</span>
									</div>
								</div>
							{/if}
						</div>
					{/each}

					<div class="sticky bottom-0 flex justify-end bg-white pt-4">
						<button
							type="submit"
							class="rounded-full bg-emerald-500 px-8 py-3 text-white transition-colors hover:bg-emerald-600 disabled:cursor-not-allowed disabled:opacity-50"
							disabled={loading || !isSubmitEnabled}
						>
							{#if loading}
								<span class="flex items-center">
									<svg
										class="-ml-1 mr-2 h-4 w-4 animate-spin"
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
									Submitting...
								</span>
							{:else}
								Submit Questionnaire
							{/if}
						</button>
					</div>
				{:else}
					<p class="text-center text-neutral-600">No questions available.</p>
				{/if}
			</form>
		</div>
	</div>
</main>

<style>
	:global(html) {
		scroll-behavior: smooth;
	}
</style>