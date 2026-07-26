<script lang="ts">
	let { eventText, selectedDay } = $props();
</script>

<div class="details">
	<h3 class="titulo">Evento:</h3>
	<p>{eventText}</p>
	<p>{selectedDay.date}</p>
</div>

<style>
	.titulo {
		width: 100%;
		text-align: center;
	}
</style>
