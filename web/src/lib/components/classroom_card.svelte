<script lang="ts">
	export let room: Classroom;
</script>

<div class="card w-84 max-w-md bg-primary">
	<div class="card-body">
		<h2 class="card-title normal-case text-white font-normal">{room.name}</h2>
		<span class="text-gray-300">{room.section}</span>
	</div>
</div>