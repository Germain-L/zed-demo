<script lang="ts">
	import { onMount } from 'svelte';

	let todos = $state<
		{ id: number; title: string; completed: boolean }[]
	>([]);
	let newTodo = $state('');

	async function fetchTodos() {
		const res = await fetch('http://localhost:8080/todos');
		todos = await res.json();
	}

	async function addTodo() {
		const res = await fetch('http://localhost:8080/todos', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({ title: newTodo, completed: false })
		});
		const created = await res.json();
		todos = [...todos, created];
		newTodo = '';
	}

	async function updateTodo(todo) {
		const res = await fetch(`http://localhost:8080/todos/${todo.id}`, {
			method: 'PUT',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify(todo)
		});
		const updated = await res.json();
		todos = todos.map((t) => (t.id === updated.id ? updated : t));
	}

	async function deleteTodo(id) {
		await fetch(`http://localhost:8080/todos/${id}`, {
			method: 'DELETE'
		});
		todos = todos.filter((t) => t.id !== id);
	}

	onMount(fetchTodos);
</script>

<h1>Todo App</h1>

<input bind:value={newTodo} placeholder="New todo" />
<button on:click={addTodo}>Add</button>

<ul>
	{#each todos as todo (todo.id)}
		<li>
			<input
				type="checkbox"
				bind:checked={todo.completed}
				on:change={() => updateTodo(todo)}
			/>
			<span style:text-decoration={todo.completed ? 'line-through' : 'none'}>
				{todo.title}
			</span>
			<button on:click={() => deleteTodo(todo.id)}>Delete</button>
		</li>
	{/each}
</ul>