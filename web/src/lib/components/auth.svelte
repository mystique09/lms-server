<script lang="ts">
	let loading = false;
	let errors = {
		username: '',
		password: ''
	};
	let data = {
		username: '',
		password: ''
	};

	async function submitForm() {
		// validate inputs
		if (data.username.length < 3) {
			errors.username = 'Username must be at least 3 characters long.';
		}
		if (data.password.length < 3) {
			errors.password = 'Password must be at least 3 characters long.';
		}

		// set errors to empty string after 2 seconds
		setTimeout(() => {
			errors.username = '';
			errors.password = '';
		}, 2000);
		loading = true;

		// submit form
		await fetch('/api/auth', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify(data)
		}).then((res) => {
			if (res.status === 200) {
				res.json().then((res) => {
					window.location.href = '/dashboard';
				});
			} else {
				res.json().then((res) => {
					errors = res.errors;
					loading = false;
				});
			}
		});

		// setTimeout(() => (loading = false), 1000);
	}
</script>

<svelte:head>
	<title>Class Management - Sign In</title>
</svelte:head>

<div class="h-1/2 m-auto flex flex-center items-center justify-center">
	<form
		on:submit|preventDefault={submitForm}
		class="shadow-lg mt-10 h-full w-full form-control rounded max-w-xs bg-gray-200 p-3"
	>
		<h1 class="mb-8 font-light text-center text-black text-sm">Welcome Back!</h1>

		<button type="button" class="btn btn-md  w-full btn-outline text-black"
			>Continue with Google</button
		>
		<button type="button" class="mt-3 btn btn-outline w-full bg-blue-600 text-white"
			>Continue with Facebook</button
		>

		<div class="divider" />

		{#if !!errors.username.length}
			<label class="label text-xs text-red-500" for="username">{errors.username}</label>
		{/if}
		<label class="input-group py-1 mt-auto">
			<span class="bg-gray-300">
				<img class="h-6 w-6" src="/user-svgrepo-com.svg" alt="user-icon" />
			</span>
			<input
				id="username"
				type="text"
				bind:value={data.username}
				placeholder="johndoe2993"
				class:input-outline={!!errors.username.length}
				class:input-error={!!errors.username.length}
				class="input input-bordered w-full max-w-xs font-semibold text-black"
				required
			/>
		</label>

		{#if !!errors.password.length}
			<label class="label text-xs text-red-500" for="password">{errors.password}</label>
		{/if}
		<label class="input-group py-1">
			<span class="bg-gray-300">
				<img class="h-6 w-6" src="/password-svgrepo-com.svg" alt="user-icon" />
			</span>
			<input
				id="password"
				type="password"
				bind:value={data.password}
				class:input-outline={!!errors.password.length}
				class:input-error={!!errors.password.length}
				placeholder="********"
				class="input input-bordered w-full max-w-xs font-semibold text-black"
				required
			/>
		</label>

		<button type="submit" tabindex="-1" class:loading class="mt-4 btn btn-primary">Sign In</button>
		<p class="text-xs text-center m-5">
			Don't have an account?
			<span><a href="/signup" class="link link-accent">Sign Up</a></span>
		</p>
	</form>
</div>