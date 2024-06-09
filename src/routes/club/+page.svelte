<script lang="ts">
    import { onMount } from "svelte";
    import { userID } from "$lib/auth";

    let event = new Promise<{
        id: string;
        name: string;
        description: string;
        managers: { id: string; name: string }[];
    }>((resolve, reject) => {
        onMount(() => {
            const id = new URLSearchParams(document.location.search).get("id");
            console.log(id);
            if (id === null) {
                reject(new Error("No id provided"));
            }

            fetch(`/api/club/details/${id}`)
                .then((res) => res.json())
                .then((data) => {
                    console.log(data);
                    resolve(data);
                });
        });
    });

    function isManager(
        club: { managers: { id: string; name: string }[] },
        userID: string | null,
    ) {
        return club.managers.some((manager) => manager.id == userID);
    }

    let newManager: null | string = null;
    let users = new Promise<{ id: string; username: string }[]>(
        (resolve, reject) => {
            onMount(() => {
                fetch("/api/user/list")
                    .then((res) => res.json())
                    .then((data) => {
                        resolve(data);
                    });
            });
        },
    );

    function addManager() {
        console.log(newManager);
        const id = new URLSearchParams(document.location.search).get("id");

        fetch(`/api/club/addManager`, {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({
                club: parseInt(id ?? "") || null, // this is a really complicated way to convert a string to a number or null
                manager: newManager,
            }),
        })
            .then((res) => res.json())
            .then((data) => {
                console.log(data);
            });
    }
</script>

<div class="flex flex-col bg-slate-400 w-screen min-h-screen items-center">
    {#await event}
        <p>loading...</p>
    {:then value}
        <h1 class="text-5xl text-center m-5">{value.name}</h1>

        <div
            class="bg-slate-500 p-9 m-2 rounded-lg text-xl w-text aspect-[2/1]"
        >
            {value.description}
        </div>

        <div class="bg-slate-500 p-2 m-2 rounded-lg text-xl w-text text-center">
            <h2 class="text-left">managers:</h2>
            <ul>
                {#each value.managers as managers}
                    <li>{managers.name}</li>
                {/each}
            </ul>
        </div>

        {#if isManager(value, $userID)}
            <div
                class="bg-slate-500 p-2 m-2 rounded-lg text-xl w-text text-center"
            >
                <h2 class="text-left">add Manager:</h2>
                <select
                    id="club"
                    name="club"
                    bind:value={newManager}
                    on:change={(e) => console.log(newManager)}
                >
                    {#await users}
                        <option>Loading...</option>
                    {:then users}
                        {#each users as user}
                            <option value={user.id}>{user.username}</option>
                        {/each}
                    {/await}
                </select>

                <button class="bg-slate-600 p-2 m-2 rounded-lg " on:click={addManager}>add Manager</button>
            </div>
        {/if}
    {:catch error}
        <p>error: {error.message}</p>
    {/await}
</div>
