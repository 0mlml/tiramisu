<script>
    import { enhance } from '$app/forms';
    export let data;
    export let form;

    let loading = false;
    let password = '';
    let confirmPassword = '';
    let passwordVisible = false;
    let confirmPasswordVisible = false;

    $: passwordsMatch = password === confirmPassword;
    $: passwordStrength = getPasswordStrength(password);

    function getPasswordStrength(pass) {
        if (!pass) return 0;
        let strength = 0;
        if (pass.length >= 8) strength += 1;
        if (pass.match(/[A-Z]/)) strength += 1;
        if (pass.match(/[0-9]/)) strength += 1;
        if (pass.match(/[^A-Za-z0-9]/)) strength += 1;
        return strength;
    }

    function handleSubmit() {
        loading = true;
        return async ({ result }) => {
            loading = false;
        };
    }
</script>

<main class="w-full min-h-screen flex items-center justify-center bg-neutral-100 py-8">
    <form 
        method="POST" 
        use:enhance={handleSubmit}
        class="shadow-md rounded-sm w-[400px] p-8 bg-white"
    >
        <h2 class="text-center text-2xl text-neutral-950 mb-8">Create Account</h2>
        
        {#if form?.error}
            <div class="mb-4 p-3 bg-red-50 border border-red-200 rounded-lg">
                <p class="text-red-600 font-light text-sm text-center">
                    {form.error.message}
                </p>
            </div>
        {/if}

        <div class="mb-4">
            <label 
                for="name" 
                class="block font-light text-neutral-800 ml-2 mb-1"
            >
                Full Name
            </label>
            <input
                type="text"
                id="name"
                name="name"
                placeholder="Enter your full name"
                class="w-full p-2 border border-neutral-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-emerald-200 focus:border-emerald-400 transition-colors"
                required
                disabled={loading}
            />
        </div>

        <div class="mb-4">
            <label 
                for="email" 
                class="block font-light text-neutral-800 ml-2 mb-1"
            >
                Email
            </label>
            <input
                type="email"
                id="email"
                name="email"
                placeholder="Enter your email"
                class="w-full p-2 border border-neutral-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-emerald-200 focus:border-emerald-400 transition-colors"
                required
                disabled={loading}
            />
        </div>

        <div class="mb-4">
            <label 
                for="password" 
                class="block font-light text-neutral-800 ml-2 mb-1"
            >
                Password
            </label>
            <div class="relative">
                <input
                    class="w-full p-2 border border-neutral-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-emerald-200 focus:border-emerald-400 transition-colors"
                    type={passwordVisible ? "text" : "password"}
                    id="password"
                    name="password"
                    bind:value={password}
                    placeholder="Create a password"
                    required
                    disabled={loading}
                />
                <button 
                    type="button"
                    class="absolute right-2 top-1/2 -translate-y-1/2 text-neutral-500 hover:text-neutral-700"
                    on:click={() => passwordVisible = !passwordVisible}
                >
                    {#if passwordVisible}
                        👁️
                    {:else}
                        👁️‍🗨️
                    {/if}
                </button>
            </div>
            {#if password}
                <div class="mt-2">
                    <div class="flex gap-1 mb-1">
                        {#each Array(4) as _, i}
                            <div class="h-1 flex-1 rounded-full transition-colors duration-200"
                                 class:bg-red-400={passwordStrength === 1 && i === 0}
                                 class:bg-yellow-400={passwordStrength === 2 && i < 2}
                                 class:bg-emerald-400={passwordStrength >= 3 && i < passwordStrength}
                                 class:bg-neutral-200={i >= passwordStrength}
                            ></div>
                        {/each}
                    </div>
                    <p class="text-xs text-neutral-600">
                        Password should be at least 8 characters with uppercase, numbers, and symbols
                    </p>
                </div>
            {/if}
        </div>

        <div class="mb-6">
            <label 
                for="confirmPassword" 
                class="block font-light text-neutral-800 ml-2 mb-1"
            >
                Confirm Password
            </label>
            <div class="relative">
                <input
                    class="w-full p-2 border border-neutral-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-emerald-200 focus:border-emerald-400 transition-colors"
                    type={confirmPasswordVisible ? "text" : "password"}
                    id="confirmPassword"
                    name="confirmPassword"
                    bind:value={confirmPassword}
                    placeholder="Confirm your password"
                    required
                    disabled={loading}
                />
                <button 
                    type="button"
                    class="absolute right-2 top-1/2 -translate-y-1/2 text-neutral-500 hover:text-neutral-700"
                    on:click={() => confirmPasswordVisible = !confirmPasswordVisible}
                >
                    {#if confirmPasswordVisible}
                        👁️
                    {:else}
                        👁️‍🗨️
                    {/if}
                </button>
            </div>
            {#if confirmPassword && !passwordsMatch}
                <p class="text-red-500 text-sm mt-1">Passwords do not match</p>
            {/if}
        </div>

        <div class="w-full flex justify-center">
            <button
                type="submit"
                class="submit text-lg text-white rounded-full py-2 px-6 transition-all duration-200 hover:bg-emerald-500 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-emerald-400 disabled:opacity-50 disabled:cursor-not-allowed"
                disabled={loading || !passwordsMatch || passwordStrength < 2}
            >
                {#if loading}
                    <span class="flex items-center">
                        <svg class="animate-spin -ml-1 mr-2 h-4 w-4 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                            <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                            <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                        </svg>
                        Creating account...
                    </span>
                {:else}
                    Create Account
                {/if}
            </button>
        </div>

        <div class="w-full flex justify-center mt-8">
            <a 
                href="/sign-in"
                class="text-neutral-600 hover:text-emerald-600 transition-colors font-light"
            >
                Already have an account? Sign in
            </a>
        </div>
    </form>
</main>

<style>
    .submit {
        background-color: #63c19d;
    }
    
    @keyframes spin {
        to {
            transform: rotate(360deg);
        }
    }
    
    .animate-spin {
        animation: spin 1s linear infinite;
    }
</style>