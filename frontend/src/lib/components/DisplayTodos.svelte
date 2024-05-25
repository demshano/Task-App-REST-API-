<script>
    import UpdateTodo from './UpdateTodo.svelte';
    import DeleteTodo from './DeleteTodo.svelte';

    export let todos;
    export let updateTodoList;
    export let deleteTodoFromList;

    let showUpdateModal = false;
    let showDeleteModal = false;
    let todoToUpdate = null;
    let todoToDelete = null;

    function openUpdateModal(todo) {
        todoToUpdate = { ...todo };
        showUpdateModal = true;
    }

    function closeUpdateModal(updatedTodo) {
        showUpdateModal = false;
        if (updatedTodo) {
            updateTodoList(updatedTodo);
        }
        todoToUpdate = null;
    }

    function openDeleteModal(todo) {
        todoToDelete = { ...todo };
        showDeleteModal = true;
    }

    function closeDeleteModal(deletedTodoId) {
        showDeleteModal = false;
        if (deletedTodoId) {
            deleteTodoFromList(deletedTodoId);
        }
        todoToDelete = null;
    }
</script>

<style>
    .todo-container {
        display: grid;
        gap: 1rem;
    }

    @media (min-width: 640px) {
        .todo-container {
            grid-template-columns: repeat(2, minmax(0, 1fr));
        }
    }

    @media (min-width: 1024px) {
        .todo-container {
            grid-template-columns: repeat(3, minmax(0, 1fr));
        }
    }
</style>

{#if showUpdateModal}
    <UpdateTodo todo={todoToUpdate} close={closeUpdateModal} />
{/if}

{#if showDeleteModal}
    <DeleteTodo todo={todoToDelete} close={closeDeleteModal} />
{/if}

<ul class="todo-container">
    {#each todos as todo (todo.id)}
        <li class="bg-white shadow-lg rounded-lg overflow-hidden">
            <div class="px-4 py-2 flex items-center justify-between">
                <div class="flex items-center">
                    <input 
                        type="checkbox" 
                        checked={todo.done} 
                        disabled 
                        class="form-checkbox h-5 w-5 text-cyan-500 mr-3 cursor-not-allowed" 
                    />
                    <span class={todo.done ? "line-through text-green-500" : "text-gray-700"}>{todo.text}</span>
                </div>
                <div class="flex space-x-2">
                    <button on:click={() => openUpdateModal(todo)} class="bg-cyan-500 hover:bg-cyan-700 text-white px-2 py-1 rounded">Update</button>
                    <button on:click={() => openDeleteModal(todo)} class="bg-red-500 hover:bg-red-700 text-white px-2 py-1 rounded">Delete</button>
                </div>
            </div>
        </li>
    {/each}
</ul>
