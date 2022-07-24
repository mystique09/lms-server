<script context="module">
	export async function load({ session }) {
		if (!session) {
			return {
				status: 301,
				redirect: '/'
			};
		}
	}
</script>

<div>Dashboard</div>
