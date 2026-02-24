<script lang="ts">
	// No Svelte 5, usamos $state para variáveis que mudam na tela
	let tasks = $state([]);

	async function loadTasks() {
		try {
			const response = await fetch('http://localhost:8080/api/tasks');
			tasks = await response.json();
		} catch (err) {
			console.error("Erro ao buscar tarefas. O backend Go está rodando?", err);
		}
	}
</script>

<main>
	<h1>BF-Planner 🚀</h1>
	<p>Unified Productivity Hub</p>

	<button onclick={loadTasks}>
		Carregar Minhas Tarefas
	</button>

	<ul>
		{#each tasks as task}
			<li>
				<input type="checkbox" checked={task.done} />
				{task.title}
			</li>
		{/each}
	</ul>
</main>

<style>
	main { font-family: 'Segoe UI', sans-serif; padding: 2rem; max-width: 600px; }
	h1 { color: #ff3e00; margin-bottom: 0.5rem; }
	button { background: #4caf50; color: white; border: none; padding: 10px 20px; border-radius: 4px; cursor: pointer; }
</style>