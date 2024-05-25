<script>
    import axios from 'axios';
    export let addTodo;

    let text = '';
    let done = false;

    async function createTodo() {
        try {
            const response = await axios.post('http://localhost:8080/todos', { text, done });
            addTodo(response.data);  // Pass the new todo back to the parent component
            text = '';
            done = false;
        } catch (error) {
            console.error('Error creating todo:', error);
        }
    }
</script>

<form on:submit|preventDefault={createTodo} class="max-w-sm mx-auto p-4 bg-white shadow-md rounded-lg mb-8">
    <div class="mb-4">
        <label for="text" class="block font-bold mb-1">Todo:</label>
        <input id="text" type="text" class="w-full px-3 py-2 border rounded-lg text-gray-700 bg-gray-200 focus:outline-none focus:border-cyan-500" bind:value={text} required />
    </div>
    <div class="mb-4">
        <label class="inline-flex items-center">
            <input type="checkbox" class="form-checkbox h-5 w-5 text-cyan-500" bind:checked={done} />
            <span class="ml-2">Done</span>
        </label>
    </div>
    <button type="submit" class="w-full bg-cyan-500 hover:bg-cyan-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline">Add Todo</button>
</form>
