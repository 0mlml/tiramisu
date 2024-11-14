<script>
	import { page } from '$app/stores';
	import { goto } from '$app/navigation';
	import * as Icons from 'lucide-svelte';

	export let data;

	const navItems = [
		{
			href: '/admin',
			label: 'Dashboard',
			icon: Icons.LayoutDashboard
		},
		{
			href: '/admin/users',
			label: 'Users',
			icon: Icons.Users
		},
		{
			href: '/admin/questions',
			label: 'Questions',
			icon: Icons.FileQuestion
		},
		{
			href: '/admin/submissions',
			label: 'Submissions',
			icon: Icons.ClipboardList
		}
	];
</script>

<div class="min-h-screen bg-gray-100">
	<!-- Sidebar -->
	<aside class="fixed inset-y-0 left-0 w-64 bg-gray-900 text-white">
		<div class="p-6">
			<h1 class="text-2xl font-bold">Admin Panel</h1>
		</div>

		<!-- User Info -->
		<div class="border-b border-gray-800 px-6 py-4">
			<div class="flex items-center space-x-3">
				{#if data.user.picture}
					<img src={data.user.picture} alt="Profile" class="h-8 w-8 rounded-full" />
				{:else}
					<div class="flex h-8 w-8 items-center justify-center rounded-full bg-gray-700">
						<span class="text-sm">{data.user.name[0].toUpperCase()}</span>
					</div>
				{/if}
				<div>
					<p class="text-sm font-medium">{data.user.name}</p>
					<p class="text-xs text-gray-400">Administrator</p>
				</div>
			</div>
		</div>

		<!-- Navigation -->
		<nav class="mt-6">
			<ul class="space-y-2">
				{#each navItems as item}
					<li>
						<a
							href={item.href}
							class="flex items-center px-6 py-3 text-gray-300 transition-colors hover:bg-gray-800 hover:text-white {$page
								.url.pathname === item.href
								? 'bg-gray-800 text-white'
								: ''}"
						>
							<svelte:component this={item.icon} size={20} class="mr-3" />
							{item.label}
						</a>
					</li>
				{/each}
			</ul>
		</nav>
	</aside>

	<!-- Main Content -->
	<main class="ml-64 p-8">
		<slot />
	</main>
</div>
