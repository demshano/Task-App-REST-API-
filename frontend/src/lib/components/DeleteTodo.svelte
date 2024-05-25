<script>
    import axios from 'axios';
    export let todo;
    export let close;

    async function deleteTodo() {
        try {
            await axios.delete(`http://localhost:8080/todos/${todo.id}`);
            close(todo.id);  // Pass the deleted todo ID back to the parent component
        } catch (error) {
            console.error('Error deleting todo:', error);
        }
    }
</script>

<div class="fixed inset-0 flex items-center justify-center bg-gray-900 bg-opacity-50">
    <div class="bg-white p-4 rounded-lg shadow-lg">
        <h2 class="text-xl font-bold mb-4">Delete Todo</h2>
        <p>Are you sure you want to delete this todo?</p>
        <div class="flex justify-end mt-4">
            <button on:click={() => close(null)} class="bg-gray-500 text-white px-4 py-2 rounded mr-2">Cancel</button>
            <button on:click={deleteTodo} class="bg-red-500 text-white px-4 py-2 rounded">Delete</button>
        </div>
    </div>
</div>
