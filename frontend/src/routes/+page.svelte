<script>
    import CreateTodo from "../lib/components/CreateTodo.svelte";
    import DisplayTodos from "../lib/components/DisplayTodos.svelte";
    import axios from "axios";
    import { onMount } from "svelte";

    let todos = [];

async function fetchTodos() {
    try {
        const response = await axios.get('http://localhost:8080/todos');
        todos = response.data;
    } catch (error) {
        console.error('Error fetching todos:', error);
    }
}

function addTodo(newTodo) {
    todos = [...todos, newTodo];
}

function updateTodoList(updatedTodo) {
    todos = todos.map(todo => todo.id === updatedTodo.id ? updatedTodo : todo);
}

function deleteTodoFromList(id) {
    todos = todos.filter(todo => todo.id !== id);
}

onMount(fetchTodos);

</script>
<p class="flex justify-center p-2 sm:text-3xl text-2xl font-bold">Daily Task</p>



<main class="p-4">
<CreateTodo addTodo={addTodo} />
<DisplayTodos todos={todos} updateTodoList={updateTodoList} deleteTodoFromList={deleteTodoFromList} />
</main>

