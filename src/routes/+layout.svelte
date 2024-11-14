<script>
	import { linkUtil } from '$lib/linkUtil';
	import { page } from '$app/stores';
	import '../app.css';
	import Topbar from '$lib/components/TopBar.svelte';
	let { children } = $props();

	const profile = $page.data.profile || null;

	if (profile !== null){
		linkUtil['links'][0]['linkText'] = profile['name'];
	} else {
		linkUtil['links'][0]['linkText'] = "Sign in!";
	}
</script>

<div class="app">
	<Topbar data={linkUtil}/>

	<main>
		{@render children()}
	</main>

	<footer>
		<p>
			Visit <a href="https://prorecovery.fi">prorecovery.fi</a> for more details
		</p>
	</footer>
</div>

<style>
	.app {
		display: flex;
		flex-direction: column;
		min-height: 100vh;
	}

	main {
		flex: 1;
		display: flex;
		flex-direction: column;
		padding: 1rem;
		width: 100%;
		max-width: 64rem;
		margin: 0 auto;
		box-sizing: border-box;
	}

	footer {
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
		padding: 12px;
	}

	footer a {
		font-weight: bold;
	}

	@media (min-width: 480px) {
		footer {
			padding: 12px 0;
		}
	}
</style>
