<script>
	import '../app.css';
	import { browser } from '$app/environment';
	import { goto } from '$app/navigation';
	import { page } from "$app/stores";

	const linkUtil = $page.data.linkUtil || null;

	if (browser) {
		let windowWidth = window.innerWidth;
	}

	function handleResize() {
		windowWidth = window.innerWidth;
		if (windowWidth >= 1024) clicked = false;
	}
</script>

<svelte:head>
	<link rel="preconnect" href="https://fonts.googleapis.com" />
	<link rel="preconnect" href="https://fonts.gstatic.com" crossorigin />
	<link href="https://fonts.googleapis.com/css2?family=Poppins:wght@300;500;700&display=swap" rel="stylesheet" />
</svelte:head>

<svelte:window on:resize={handleResize} />

<div class="entry">
	{#each linkUtil.links as link}
		{#if link.id === 'profile'}
			<h1 class="greeting font-bold">Welcome, {link.linkText} 👋</h1>
		{:else}
			<h1 class="greeting font-bold">Welcome</h1>
		{/if}
	{/each}
    
	<form class="entry-form" action="/survey">
		<button class="entry-button continue">Start Survey</button>
	</form>
	
	<form class="entry-form" action="/logout" method="POST">
		<button type="submit" class="entry-button quit">Logout</button>
	</form>
</div>

<style>
	.greeting {
		font-family: 'Poppins', sans-serif;
		font-weight: 700;
		font-size: 2.5rem;
		color: #333;
		text-align: center;
		margin-bottom: 1.5em;
	}

	.entry {
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
		min-height: 100vh;
		background-color: #f7f9fc;
		padding: 2rem;
	}

	.entry-button {
		width: 100%;
		padding: 1rem;
		font-family: 'Poppins', sans-serif;
		font-weight: 500;
		color: #fff;
		border-radius: 8px;
		transition: transform 0.3s ease, background-color 0.3s ease;
		box-shadow: 0px 4px 12px rgba(0, 0, 0, 0.1);
	}

	.continue {
		background-color: #63c19d;
	}

	.quit {
		background-color: #d9534f;
	}

	.entry-button:hover {
		transform: scale(1.05);
	}

	.entry-button:focus {
		outline: none;
		box-shadow: 0px 0px 0px 4px rgba(0, 123, 255, 0.3);
	}

	.entry-form {
		width: 80%;
		max-width: 400px;
		margin-top: 2em;
	}

</style>
