<script>
	async function getCookie() {
		const response = await fetch('/api/v1/test', {
			method: 'POST',
			headers: {
				'Content-type': 'application/json',
				Accept: 'application/json'
			}
		});
		const data = await response.text();
		$session = 'jsjsjsjsjs';
		console.log(JSON.stringify($session));
	}
	import { session } from '$app/stores';
</script>

<h1 class="font-bold text-3xl">Welcome to SvelteKit</h1>
<p class="font-light text-lg">
	Visit <a href="https://kit.svelte.dev">kit.svelte.dev</a> to read the documentation
</p>

<h2>Class Management</h2>
<input type="button" on:click={getCookie} value="Get cookie" />
{$session}
