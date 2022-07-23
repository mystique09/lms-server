<script lang="ts">
	let toggle = false;

	const toggleNav = () => {
		toggle = !toggle;
	};
</script>

<nav class="container m-auto navbar flex flex-col items-end bg-gray-400/40">
	<div class="flex flex-row w-full">
		<div class="flex-1">
			<a href="/" class="btn btn-ghost text-base normal-case">Class Manager</a>
		</div>

		<a
			href="https://github.com/mystique09"
			rel="noopener"
			class="btn mr-4 btn-sm bg-gray-300 hover:bg-gray-400 btn-circle"
		>
			<img src="/github-svgrepo-com.svg" alt="github-icon" />
		</a>

		<div class="flex-none">
			<button
				name="menu-btn"
				on:click={() => (toggle = !toggle)}
				class="btn btn-md btn-secondary text-white"
			>
				<svg
					xmlns="http://www.w3.org/2000/svg"
					fill="none"
					viewBox="0 0 24 24"
					class="inline-block w-5 h-5 stroke-current"
					><path
						stroke-linecap="round"
						stroke-linejoin="round"
						stroke-width="2"
						d="M4 6h16M4 12h16M4 18h16"
					/></svg
				>
			</button>
		</div>
	</div>
	<div class="menu w-full flex flex-col items-end" class:hidden={!toggle}>
		<div class="divider" />
		<menu>
			<li>
				<a on:click={toggleNav} class="btn btn-ghost btn-md text-accent" href="/auth">Login</a>
			</li>
			<li>
				<a on:click={toggleNav} class="btn btn-ghost btn-md text-accent" href="/signup">Signup</a>
			</li>
			<li>
				<a on:click={toggleNav} class="btn btn-ghost btn-md text-accent" href="/#about">About Us</a>
			</li>
		</menu>
	</div>
</nav>