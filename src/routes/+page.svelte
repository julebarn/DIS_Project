<script lang="ts">
    import { onMount } from "svelte";
    import EventLi from "../lib/components/event_li.svelte";
    import { datetimeFormater } from "$lib/datetimeFormater";
    import { userID, refreshTokens } from "$lib/auth";

    let events = new Promise<any[]>((resolve, reject) => {});
    let clubs = new Promise<any[]>((resolve, reject) => {});

    onMount(() => {
        events = fetch("/api/event/future")
            .then((res) => res.json())
            .then((data) => {
                console.log(data);
                return data;
            });

        clubs = fetch("/api/club/list")
            .then((res) => res.json())
            .then((data) => {
                console.log(data);
                return data;
            });
    });

    function arrN(n: number) {
        return Array.from({ length: n }, (_, i) => ({
            ID: i,
            Name: "club " + i,
        }));
    }
</script>

<div class="flex flex-col items-stretch bg-slate-400 w-screen min-h-screen">
    <h1 class="text-5xl text-center m-5">what is happening @ DIKU</h1>

    <div class="flex justify-center">
        {#if $userID == null}
            <a class="bg-slate-500 p-2 rounded-lg m-2" href="login">Login</a>
            <a class="bg-slate-500 p-2 rounded-lg m-2" href="register"
                >register</a
            >
        {:else}
            <a class="bg-slate-500 p-2 rounded-lg m-2" href="/Logout">Logout</a>
            <a class="bg-slate-500 p-2 rounded-lg m-2" href="/create/event"
                >Create event</a
            >
            <a class="bg-slate-500 p-2 rounded-lg m-2" href="/create/club"
                >Create club</a
            >
        {/if}
    </div>

    <div class="flex justify-center">
        {#await events}
            <p>loading...</p>
        {:then value}
            {#if value}
                <ul class="flex flex-col items-stretch w-text">
                    {#each value as event}
                        <li>
                            <EventLi
                                name={event.name}
                                start_time={datetimeFormater(event.start_time)}
                                ID={event.id}
                            />
                        </li>
                    {/each}
                </ul>
            {:else}
                <p>no events</p>
            {/if}
        {:catch error}
            <p>error: {error.message}</p>
        {/await}
    </div>
    <h2 class="text-3xl text-center m-5">Clubs@DIKU</h2>
    <div class="flex justify-center">
        {#await clubs}
            <p>loading...</p>
        {:then value}
            {#if value}
                <ul class="flex flex-col items-stretch w-text">
                    {#each value as club}
                        <li>
                            <a class="bg-slate-500 p-9 m-2 rounded-lg flex text-xl" href="/club?id={club.ID}">{club.Name}</a>
                        </li>
                    {/each}
                </ul>
            {:else}
                <p>no clubs</p>
            {/if}
        {:catch error}
            <p>error: {error.message}</p>
        {/await}
    </div>
</div>
