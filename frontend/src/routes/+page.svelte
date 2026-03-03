<script lang="ts">
	let tasks = $state([]);
	let taskTitle = $state("");
  	let message = $state("");

	async function loadTasks() {
		try {
			const response = await fetch('https://crispy-broccoli-gpg6rxxrvxwcw74-8080.app.github.dev/api/tasks');
			tasks = await response.json();
		} catch (err) {
			console.error("Erro ao buscar tarefas. O backend Go está rodando?", err);
		}
	}

	async function sendTask() {
    	const data = { StrTitle: taskTitle };

		try {
			const response = await fetch("https://crispy-broccoli-gpg6rxxrvxwcw74-8080.app.github.dev/api/tasks", {
				method: "POST",
				headers: {
				"Content-Type": "application/json",
				},
				body: JSON.stringify(data), 
			});

			if (response.ok) {
				message = "Tarefa enviada com sucesso!";
				taskTitle = ""; // Limpa o campo
			} else {
				message = "Erro ao enviar: " + response.statusText;
			}
		} catch (error) {
			message = "Erro de conexão: " + error.message;
		}
		loadTasks();
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
	<input type="text" id="oi" bind:value={taskTitle}>
	<br>
	<button onclick={sendTask}>
		ENVIAR Minhas Tarefas
	</button>
	<p>{message}</p>
</main>

<style>
	main { font-family: 'Segoe UI', sans-serif; padding: 2rem; max-width: 600px; }
	h1 { color: #ff3e00; margin-bottom: 0.5rem; }
	button { background: #4caf50; color: white; border: none; padding: 10px 20px; border-radius: 4px; cursor: pointer; }
</style>