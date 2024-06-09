<script lang="ts">
    import { onMount } from "svelte";

    let event = new Promise<
        {
            organizers: { name: string }[];
            club_id: number;
            name: string;
            description: string;
            start_time: string;
            end_time: string;
            place: string;
            oneday: boolean;
            date: string;
        }[]
    >((resolve, reject) => {
        onMount(() => {
            const id = new URLSearchParams(document.location.search).get("id");
            console.log(id);
            if (id === null) {
                reject(new Error("No id provided"));
            }
            const event = fetch(`/api/event/details/${id}`)
                .then((res) => res.json())
                .then((data) => {
                    console.log(data);

                    let start_time = new Date(data.start_time);
                    let end_time = new Date(data.end_time);

                    if (start_time.toDateString() === end_time.toDateString()) {
                        data.oneday = true;
                        data.start_time = `${start_time.getHours()}:${start_time.getMinutes()}`;
                        data.end_time = `${end_time.getHours()}:${end_time.getMinutes()}`;
                        data.date = `${start_time.getDay()}/${start_time.getMonth()}/${start_time.getFullYear()}`;
                    } else {
                        data.oneday = false;
                        data.start_time = `${start_time.getHours()}:${start_time.getMinutes()}`;
                        data.end_time = `${end_time.getHours()}:${end_time.getMinutes()}`;
                        data.end_date = `${end_time.getDay()}/${end_time.getMonth()}/${end_time.getFullYear()}`;
                        data.start_date = `${start_time.getDay()}/${start_time.getMonth()}/${start_time.getFullYear()}`;
                    }

                    return data;
                });

            resolve(event);
        });
    });
</script>

<div class="flex flex-col items-stretch bg-slate-400 w-screen min-h-screen">
    {#await event}
        <p>loading...</p>
    {:then value}
        <h1 class="text-5xl text-center m-5">{value.name}</h1>
        <div class="flex flex-col items-center">
            <div class="bg-slate-500 p-9 m-2 rounded-lg text-xl w-text flex">
                {#if value.oneday}
                    <div class="grid grid-cols-3 w-full">
                        <span>From: {value.start_time} </span>
                        <span>To: {value.end_time} </span>
                        <span>On: {value.date} </span>
                    </div>
                {:else}
                    <div class="grid grid-cols-2 w-full">
                        <div>
                            <span>From: {value.start_time}</span><br />
                            <span>On: {value.start_date}</span>
                        </div>
                        <div>
                            <span>To: {value.end_time}</span><br />
                            <span>On: {value.end_date}</span>
                        </div>
                    </div>
                {/if}
            </div>

            <div
                class="bg-slate-500 p-9 m-2 rounded-lg text-xl w-text grid grid-cols-2"
            >
                Location: {value.place}
            </div>

            <div
                class="bg-slate-500 p-9 m-2 rounded-lg text-xl w-text aspect-[2/1]"
            >
                {value.description}
            </div>

            {#if value.club_id != null}
                <div
                    class="bg-slate-500 p-2 m-2 rounded-lg text-xl w-text text-center"
                >
                    <a href="/club?id={value.club_id}">Go to club</a>
                </div>
            {/if}

            <div class="bg-slate-500 p-2 m-2 rounded-lg text-xl w-text text-center">
                <h2 class="text-left">Organizers : </h2>
                <ul>
                    {#each value.organizers as organizer}
                        <li>{organizer.name}</li>
                    {/each}
                </ul>
            </div>
        </div>
    {:catch error}
        <p>error: {error.message}</p>
    {/await}
</div>
