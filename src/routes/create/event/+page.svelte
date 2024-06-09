<script lang="ts">
    import { goto } from "$app/navigation";
    import { onMount } from "svelte";
    import { userID } from "$lib/auth";
    import { writable, derived } from "svelte/store";
    let name = "";
    let description = "";
    let place = "";
    let dateTimeStart = "";
    let dateTimeEnd = "";
    let club: null | string = null;
    // TODO add club

    function createEvent() {
        if (name == "") {
            alert("Name is required");
            return;
        }
        if (description == "") {
            alert("Description is required");
            return;
        }
        if (place == "") {
            alert("Place is required");
            return;
        }
        if (dateTimeStart == "") {
            alert("Start date is required");
            return;
        }
        if (dateTimeEnd == "") {
            alert("End date is required");
            return;
        }

        const body = JSON.stringify({
            name: name,
            description: description,
            place: place,
            start_time: dateTimeStart,
            end_time: dateTimeEnd,
            club_id: club,
        });
        console.log(body);

        fetch("/api/event/create", {
            method: "POST",
            credentials: "include",
            headers: {
                "Content-Type": "application/json",
            },
            body: body,
        })
            .then((r) => {
                console.log("Success:", r);
                goto("/");
            })
            .catch((error) => {
                console.error("Error:", error);
            });
    }

    let clubs = new Promise<any[]>((resolve, reject) => {});
    let users = writable<{ id: number; name: string }[]>([]);

    onMount(async () => {
        clubs = fetch("/api/club/isOrganizer")
            .then((res) => res.json())
            .then((data) => {
                console.log(data);
                return data;
            });
        const fetchedUsers = await fetchUsers();

        users.set(fetchedUsers);

        if (fetchedUsers.length > 0) {
            newOrganizer = fetchedUsers[0];
        }
    });

    let organizers = writable<{ id: number; name: string }[]>([]);
    organizers.subscribe((value) => {
        console.log("organizers changed", value);
    });

    let newOrganizer: { id: number; name: string } | null = null;

    function fetchUsers(): Promise<{ id: number; name: string }[]> {
        return new Promise<{ id: number; name: string }[]>((resolve, reject) =>
            fetch("/api/user/list")
                .then((res) => res.json())
                .then((data: { id: number; username: string }[]) => {
                    console.log(data);
                    let Data = data.map((user) => ({
                        id: user.id,
                        name: user.username,
                    }));

                    let user = Data.find(
                        (user) => user.id.toString() == $userID,
                    );
                    if (user != undefined) {
                        organizers.update((value) => {
                            console.log("adding user", user);
                            return [...value, user];
                        });
                    }

                    resolve(Data);
                }),
        );
    }

    function validateDateTime() {

        const timeStart = new Date(dateTimeStart);
        const timeEnd = new Date(dateTimeEnd);

        if (Date.parse(dateTimeStart) >= Date.parse(dateTimeEnd)) {
            dateTimeStart = "";
            dateTimeEnd = "";
            alert("End date must be after start date");
        }
    }

    //let users = new Promise<{ id: number; name: string }[]>(
    //    (resolve, reject) => {
    //        onMount(() => {
    //            fetch("/api/user/list")
    //                .then((res) => res.json())
    //                .then((data: { id: number; username: string }[]) => {
    //                    console.log(data);
    //                    let Data = data.map((user) => ({
    //                        id: user.id,
    //                        name: user.username,
    //                    }));

    //                    let user = Data.find(
    //                        (user) => user.id.toString() == $userID,
    //                    );
    //                    if (user != undefined) {
    //                        organizers.update((value) => {
    //                            console.log("adding user", user);
    //                            return [...value, user];
    //                        });
    //                    }

    //                    resolve(Data);
    //                });
    //        });
    //    },
    //);

    function addOrganizer() {
        organizers.update((value) => {
            console.log("adding organizer", newOrganizer);
            if (newOrganizer == null) {
                return value;
            }

            // Check if the user is already in the list, if it is, do nothing
            if (value.find((user) => user.id == newOrganizer!.id)) {
                return value;
            }

            // Otherwise add the user to the list
            return [...value, { id: newOrganizer.id, name: newOrganizer.name }];
        });
    }

</script>

<div class="flex flex-col items-stretch bg-slate-400 w-screen min-h-screen">
    <h1 class="text-5xl text-center m-5">Create Event</h1>
    <div class="flex flex-col justify-center items-center space-y-1">
        <label for="name" class="text-xl">Name </label>
        <input
            class="bg-slate-500"
            type="text"
            id="name"
            name="name"
            bind:value={name}
        />
        <br />
        <label for="description" class ="text-xl">Description </label>
        <input
            class="bg-slate-500"
            type="text"
            id="description"
            name="description"
            bind:value={description}
        />
        <br />
        <label for="place" class ="text-xl">Place </label>
        <input
            class="bg-slate-500"
            type="text"
            id="place"
            name="place"
            bind:value={place}
        />
        <br />
        <label for="dateTimeStart" class ="text-xl">Start Date </label>
        <input
            class="bg-slate-500"
            type="datetime-local"
            id="dateTimeStart"
            name="dateTimeStart"
            bind:value={dateTimeStart}
            on:change={validateDateTime}
        />
        <br />
        <label for="dateTimeEnd" class ="text-xl">End Date </label>
        <input
            class="bg-slate-500"
            type="datetime-local"
            id="dateTimeEnd"
            name="dateTimeEnd"
            bind:value={dateTimeEnd}
            on:change={validateDateTime}
        />
        <br />

        <label for="club" class ="text-xl">Club</label>
        <select
            id="club"
            name="club"
            bind:value={club}
            on:change={(e) => console.log(club)}
        >
            <option value="null"> no club </option>
            {#await clubs}
                <option>Loading...</option>
            {:then clubs}
                {#each clubs as club}
                    <option value={club.ID}>{club.Name}</option>
                {/each}
            {/await}
        </select>
        <br />
        <h2 class="text-3xl text-center m-5">Organizers</h2>
        <ul>
            {#each $organizers as organizer}
                <li data-id={organizer.id}>{organizer.name}</li>
            {/each}
        </ul>
        {#if $users.length === 0}
            <select disabled>
                <option>Loading...</option>
            </select>
        {:else}
            <select
                id="club"
                name="club"
                bind:value={newOrganizer}
                on:change={(e) => console.log(newOrganizer)}
            >
                {#each $users as user}
                    {#if !$organizers.includes(user)}
                        <option value={user}>{user.name}</option>
                    {/if}
                {/each}
            </select>
        {/if}
        <button on:click={addOrganizer}>Add Manager</button>

        <br />

        <a href="javascript:void(0)" class="bg-slate-500 p-2 m-6 rounded-lg text-xl text-center w-64" on:click={createEvent}>
            Create Event
        </a>
    </div>
</div>
