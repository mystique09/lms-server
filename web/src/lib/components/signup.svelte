<script lang="ts">
	let loading = false;
	let errors = {
		username: '',
		email: '',
		password: ''
	};
	let data = {
		username: '',
		email: '',
		password: ''
	};

	async function submitForm() {
		loading = true;

		// validate inputs
		if (data.username.length < 8) {
			errors.username = 'Minimum length of 8 characters';
		}

		// validate email
		if (!/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(data.email)) {
			errors.email = 'Invalid email format.';
		}

		// validate password length
		if (data.password.length < 8) {
			errors.password = 'Minimum length of 8 alphanumeric characters';
		}

		 const req = await fetch('/api/auth/signup', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify(data)
		});
		const res = await req.json();

		if( res.status === 200) {
			window.location.href = '/dashboard';
			return;
		} else {
			errors = res.errors;
			loading = false;

			 		// set back to empty string after 2 seconds
		setTimeout(() => {
			errors.username = '';
			errors.email = '';
			errors.password = '';
		}, 2000);

		setTimeout(() => (loading = false), 1000);	
			return;
		}
	}
</script>

<svelte:head>
	<title>Class Management - Sign Up</title>
</svelte:head>

<div class="h-1/2 m-auto flex flex-center items-center justify-center">
	<form
		on:submit|preventDefault={submitForm}
		class="my-10 h-full w-full form-control rounded max-w-xs bg-gray-100 shadow-md p-3"
	>
		<h1 class="mb-8 font-light text-center text-black text-sm">Create new Account</h1>

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
				class="input input-bordered w-full max-w-xs"
				required
			/>
		</label>

		{#if !!errors.email.length}
			<label class="label text-xs text-red-500" for="username">{errors.email}</label>
		{/if}
		<label class="input-group py-1">
			<span class="bg-gray-300">
				<img class="h-6 w-6" src="/email-svgrepo-com.svg" alt="user-icon" />
			</span>
			<input
				id="email"
				bind:value={data.email}
				type="email"
				placeholder="johndoe2993@email.com"
				class:input-outline={!!errors.email.length}
				class:input-error={!!errors.email.length}
				class="input input-bordered w-full max-w-xs"
				required
			/>
		</label>

		{#if !!errors.password.length}
			<label class="label text-xs text-red-500" for="username">{errors.password}</label>
		{/if}
		<label class="input-group py-1">
			<span class="bg-gray-300">
				<img class="h-6 w-6" src="/password-svgrepo-com.svg" alt="user-icon" />
			</span>
			<input
				id="password"
				type="password"
				bind:value={data.password}
				placeholder="********"
				class:input-outline={!!errors.password.length}
				class:input-error={!!errors.password.length}
				class="input input-bordered w-full max-w-xs"
				required
			/>
		</label>

		<button type="submit" tabindex="-1" class:loading class="mt-4 btn btn-primary">Sign Up</button>
		<p class="text-xs text-center m-5">
			Already have an account?
			<span><a href="/auth" class="link link-accent">Login</a></span>
		</p>
	</form>
</div>