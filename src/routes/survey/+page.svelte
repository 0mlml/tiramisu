<script>
	import { enhance } from '$app/forms';
	import { goto } from '$app/navigation';

	export let data;
	export let form;

	let loading = false;
	let currentQuestion = 0;
	let answers = new Map();
	let currentAnswer = undefined;

	$: questions = data.questions || [];
	$: question = questions[currentQuestion];
	$: isLastQuestion = currentQuestion === questions.length - 1;
	$: isFirstQuestion = currentQuestion === 0;
	$: progress = ((currentQuestion + 1) / questions.length) * 100;

	function handleAnswerChange(value) {
		currentAnswer = value;
		answers.set(question.id, value);
		answers = answers; // Trigger reactivity
	}

	function handleNext() {
		if (currentQuestion < questions.length - 1) {
			currentQuestion++;
			currentAnswer = answers.get(questions[currentQuestion].id);
		}
	}

	function handlePrevious() {
		if (currentQuestion > 0) {
			currentQuestion--;
			currentAnswer = answers.get(questions[currentQuestion].id);
		}
	}

	function handleSubmit() {
		loading = true;
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

	// Set initial answer if it exists
	$: if (question) {
		currentAnswer = answers.get(question.id);
	}
</script>

<main class="min-h-screen w-full bg-neutral-100 px-4 py-8">
	<div class="mx-auto max-w-2xl">
		<!-- Progress bar -->
		<div class="mb-8 h-2.5 w-full rounded-full bg-neutral-200">
			<div
				class="h-2.5 rounded-full bg-emerald-500 transition-all duration-300"
				style="width: {progress}%"
			></div>
		</div>

		<div class="rounded-lg bg-white p-8 shadow-md">
			{#if form?.error}
				<div class="mb-6 rounded-lg border border-red-200 bg-red-50 p-4">
					<p class="text-center text-red-600">{form.error}</p>
				</div>
			{/if}

			<form method="POST" action="?/submit" use:enhance={handleSubmit}>
				{#if question}
					<div class="mb-8">
						<h2 class="mb-6 text-xl font-medium">
							{question.question}
						</h2>

						{#if question.type === 'scale'}
							<div class="flex flex-col items-center">
								<div class="mb-2 flex w-full max-w-md justify-between">
									{#each getScaleLabels(question.min, question.max) as value}
										<label class="group flex cursor-pointer flex-col items-center">
											<input
												type="radio"
												name="question_{question.id}"
												{value}
												checked={currentAnswer === value}
												class="peer sr-only"
												on:change={() => handleAnswerChange(value)}
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

					<div class="mt-8 flex justify-between">
						<button
							type="button"
							class="rounded-full px-6 py-2 text-neutral-600 transition-colors hover:bg-neutral-100"
							disabled={isFirstQuestion || loading}
							on:click={handlePrevious}
						>
							Previous
						</button>

						{#if isLastQuestion}
							<button
								type="submit"
								class="rounded-full bg-emerald-500 px-6 py-2 text-white transition-colors hover:bg-emerald-600 disabled:cursor-not-allowed disabled:opacity-50"
								disabled={loading || !answers.has(question.id)}
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
									Submit
								{/if}
							</button>
						{:else}
							<button
								type="button"
								class="rounded-full bg-emerald-500 px-6 py-2 text-white transition-colors hover:bg-emerald-600 disabled:cursor-not-allowed disabled:opacity-50"
								disabled={!answers.has(question.id)}
								on:click={handleNext}
							>
								Next
							</button>
						{/if}
					</div>
				{:else}
					<p class="text-center text-neutral-600">No questions available.</p>
				{/if}
			</form>
		</div>
	</div>
</main>
