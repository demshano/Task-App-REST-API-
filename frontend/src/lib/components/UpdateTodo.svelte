<script>
    import axios from 'axios';
    export let todo;
    export let close;

    let newText = todo.text;
    let newDone = todo.done;

    async function updateTodo() {
        try {
            const updatedTodo = { ...todo, text: newText, done: newDone };
            const response = await axios.put(`http://localhost:8080/todos/${todo.id}`, updatedTodo);
            close(response.data);  // Pass the updated todo back to the parent component
        } catch (error) {
            console.error('Error updating todo:', error);
        }
    }
</script>

<div class="fixed inset-0 flex items-center justify-center bg-gray-900 bg-opacity-50">
    <div class="bg-white p-4 rounded-lg shadow-lg">
        <h2 class="text-xl font-bold mb-4">Update Todo</h2>
        <input type="text" bind:value={newText} class="border p-2 w-full mb-4" />
        <input type="checkbox" bind:checked={newDone} class="form-checkbox h-5 w-5 text-cyan-500 mb-4" />
        <div class="flex justify-end">
            <button on:click={() => close(null)} class="bg-gray-500 text-white px-4 py-2 rounded mr-2">Cancel</button>
            <button on:click={updateTodo} class="bg-cyan-500 text-white px-4 py-2 rounded">Update</button>
        </div>
    </div>
</div>
