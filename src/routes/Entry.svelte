<script>
    import  "../app.css";
	import {browser} from '$app/environment';
    export let data;
    
    if (browser) {
        let windowWidth = window.innerWidth;
    }

    function handleResize(){
            windowWidth = window.innerWidth;
            if(windowWidth >= 1024) clicked = false;
    }

</script>

<svelte:window on:resize={handleResize} />
<div class='entry'>
    {#each data.links as link}
        {#if link.id === 'profile'}
        <h1 class="font-bold">Welcome, {link.linkText}!</h1>
        {:else}
        <h1 class="font-bold">Welcome</h1>
        {/if}
    {/each}
    <button class="entry-button continue text-white ring-2 rounded-md ring-green-800 transition-duration-300 transform hover:scale-105 hover:bg-green-800 focus:outline-none focus:ring-4 focus:ring-green-500 focus:ring-offset-2">Start survey</button>
	<button class="entry-button quit text-white ring-2 rounded-md ring-red-800 transition-duration-300 transform hover:scale-105 hover:bg-red-800 focus:outline-none focus:ring-4 focus:ring-red-500 focus:ring-offset-2">Quit</button>
</div>

<style>
    .entry {
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
		flex: 2;
		width: 100%;
	}

	.entry-button{
		width: 80%;
		margin-top: 5em;
		padding: 1em;
	}

	.continue {
		background-color: #08af48;
	}

	.quit {
		background-color: #FF605C;
	}
</style>