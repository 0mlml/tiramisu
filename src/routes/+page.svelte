<script>
	import '../app.css';
	import { browser } from '$app/environment';
	import { linkUtil } from '$lib/linkUtil';
	import { goto } from '$app/navigation';

	export let data;

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
	<link
		href="https://fonts.googleapis.com/css2?family=Rubik:ital,wght@0,300..900;1,300..900&display=swap"
		rel="stylesheet"
	/>
</svelte:head>

<svelte:window on:resize={handleResize} />
<div class="entry">
	{#each linkUtil.links as link}
		{#if link.id === 'profile'}
			<h1 class="rubik-a font-bold">Welcome, {link.linkText}👋</h1>
		{:else}
			<h1 class="font-bold">Welcome</h1>
		{/if}
	{/each}
    <form class="entry-form" action="/survey">
        <button
		    class="entry-button rubik-a continue transition-duration-300 transform rounded-md text-white ring-4 ring-gray-600 hover:scale-105 hover:bg-green-800 focus:outline-none focus:ring-4 focus:ring-green-500 focus:ring-offset-2"
		>
            Start survey
        </button>
    </form>
	<form class="entry-form" action="/logout" method="POST">
		<button
			type="submit"
			class="entry-button rubik-a quit transition-duration-300 transform rounded-md text-white ring-4 ring-gray-600 hover:scale-105 hover:bg-red-800 focus:outline-none focus:ring-4 focus:ring-red-500 focus:ring-offset-2"
		>
			Logout
		</button>
	</form>
</div>

<style>
	.rubik-a {
		font-family: 'Rubik', sans-serif;
		font-optical-sizing: auto;
		font-weight: 500;
		font-style: normal;
	}

	.entry {
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
		flex: 2;
		width: 100%;
	}

    .entry-button{
        width: 100%;
		padding: 1em;
    }

	.entry-form {
		width: 80%;
		margin-top: 5em;
		padding: 1em;
	}

	.continue {
		background-color: #63c19d;
	}

	.quit {
		background-color: #b91c1c;
	}
</style>
