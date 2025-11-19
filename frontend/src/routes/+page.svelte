<script lang="ts">
	import { onMount } from 'svelte';
	import { fade, fly, slide } from 'svelte/transition';
	import { quintOut } from 'svelte/easing';

	let todos = $state<
		{ id: number; title: string; completed: boolean }[]
	>([]);
	let newTodo = $state('');
	let inputElement: HTMLInputElement;

	async function fetchTodos() {
		try {
			const res = await fetch(`${import.meta.env.VITE_API_URL}/todos`);
			if (res.ok) {
				todos = await res.json();
			}
		} catch (e) {
			console.error('Failed to fetch todos', e);
		}
	}

	async function addTodo() {
		if (!newTodo.trim()) return;
		
		try {
			const res = await fetch(`${import.meta.env.VITE_API_URL}/todos`, {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify({ title: newTodo, completed: false })
			});
			const created = await res.json();
			todos = [...todos, created];
			newTodo = '';
		} catch (e) {
			console.error('Failed to add todo', e);
		}
	}

	async function updateTodo(todo: { id: number; title: string; completed: boolean }) {
		// Optimistic update
		const originalTodos = [...todos];
		todos = todos.map((t) => (t.id === todo.id ? todo : t));

		try {
			const res = await fetch(`${import.meta.env.VITE_API_URL}/todos/${todo.id}`, {
				method: 'PUT',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify(todo)
			});
			if (!res.ok) {
				todos = originalTodos; // Revert on failure
			}
		} catch (e) {
			todos = originalTodos;
			console.error('Failed to update todo', e);
		}
	}

	async function deleteTodo(id: number) {
		// Optimistic update
		const originalTodos = [...todos];
		todos = todos.filter((t) => t.id !== id);

		try {
			const res = await fetch(`${import.meta.env.VITE_API_URL}/todos/${id}`, {
				method: 'DELETE'
			});
			if (!res.ok) {
				todos = originalTodos;
			}
		} catch (e) {
			todos = originalTodos;
			console.error('Failed to delete todo', e);
		}
	}

	function handleKeydown(e: KeyboardEvent) {
		if (e.key === 'Enter') {
			addTodo();
		}
	}

	onMount(fetchTodos);
</script>

<div class="min-h-screen w-full bg-gradient-to-br from-indigo-500 via-purple-500 to-pink-500 flex items-center justify-center p-4 sm:p-8">
	<div class="w-full max-w-md bg-white/10 backdrop-blur-xl border border-white/20 rounded-3xl shadow-2xl overflow-hidden flex flex-col max-h-[80vh]">
		
		<!-- Header -->
		<div class="p-6 pb-4 border-b border-white/10">
			<h1 class="text-3xl font-bold text-white mb-1 tracking-tight">Tasks</h1>
			<p class="text-white/60 text-sm font-medium">
				{todos.filter(t => !t.completed).length} items remaining
			</p>
		</div>

		<!-- List -->
		<div class="flex-1 overflow-y-auto p-4 space-y-2 custom-scrollbar">
			{#if todos.length === 0}
				<div class="text-center py-10 text-white/40" in:fade>
					<p class="text-lg">No tasks yet</p>
					<p class="text-sm">Add one below to get started</p>
				</div>
			{:else}
				{#each todos as todo (todo.id)}
					<div
						class="group flex items-center gap-3 p-3 rounded-xl bg-white/5 hover:bg-white/10 transition-all duration-200 border border-transparent hover:border-white/10"
						in:fly={{ y: 20, duration: 300 }}
						out:slide={{ duration: 200 }}
					>
						<button
							class="relative flex-shrink-0 w-6 h-6 rounded-full border-2 flex items-center justify-center transition-colors duration-200
							{todo.completed ? 'bg-green-400 border-green-400' : 'border-white/40 hover:border-white/60'}"
							onclick={() => updateTodo({ ...todo, completed: !todo.completed })}
							aria-label={todo.completed ? "Mark as incomplete" : "Mark as complete"}
						>
							{#if todo.completed}
								<svg class="w-3.5 h-3.5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
									<path stroke-linecap="round" stroke-linejoin="round" stroke-width="3" d="M5 13l4 4L19 7" />
								</svg>
							{/if}
						</button>

						<span 
							class="flex-1 text-white/90 transition-all duration-200 truncate
							{todo.completed ? 'line-through text-white/40' : ''}"
						>
							{todo.title}
						</span>

						<button
							class="opacity-0 group-hover:opacity-100 p-2 text-white/40 hover:text-red-400 hover:bg-white/10 rounded-lg transition-all duration-200 focus:opacity-100"
							onclick={() => deleteTodo(todo.id)}
							aria-label="Delete todo"
						>
							<svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
								<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
							</svg>
						</button>
					</div>
				{/each}
			{/if}
		</div>

		<!-- Input Area -->
		<div class="p-4 bg-white/5 border-t border-white/10 backdrop-blur-md">
			<div class="relative flex items-center">
				<input
					bind:this={inputElement}
					bind:value={newTodo}
					onkeydown={handleKeydown}
					type="text"
					placeholder="Add a new task..."
					class="w-full bg-white/10 text-white placeholder-white/40 rounded-xl py-3.5 pl-4 pr-12 border border-white/10 focus:outline-none focus:ring-2 focus:ring-purple-400/50 focus:bg-white/20 transition-all"
				/>
				<button
					onclick={addTodo}
					class="absolute right-2 p-2 bg-purple-500 hover:bg-purple-400 text-white rounded-lg transition-colors shadow-lg disabled:opacity-50 disabled:cursor-not-allowed"
					disabled={!newTodo.trim()}
					aria-label="Add todo"
				>
					<svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
						<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
					</svg>
				</button>
			</div>
		</div>
	</div>
</div>

<style>
	/* Custom scrollbar for Webkit browsers */
	.custom-scrollbar::-webkit-scrollbar {
		width: 6px;
	}
	.custom-scrollbar::-webkit-scrollbar-track {
		background: transparent;
	}
	.custom-scrollbar::-webkit-scrollbar-thumb {
		background-color: rgba(255, 255, 255, 0.1);
		border-radius: 20px;
	}
	.custom-scrollbar::-webkit-scrollbar-thumb:hover {
		background-color: rgba(255, 255, 255, 0.2);
	}
</style>