<script lang="ts">
    import { goto } from "$app/navigation";

    let name = "";
    let description = "";

    function createClub() {

        if (name === "" || description === "") {
            alert("Please fill in all fields");
            return;
        }


        const body = JSON.stringify({
            name: name,
            description: description,
        });
        console.log(body);

        fetch("/api/club/create", {
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
</script>

<div class="flex flex-col items-stretch bg-slate-400 w-screen min-h-screen">
    <h1 class="text-5xl text-center m-5">Create Club</h1>

    <div class="flex flex-col justify-center items-center space-y-1">
        <label for="name" class="text-xl">Name</label>
        <input
            type="text"
            id="name"
            name="name"
            class="bg-slate-500 w-64"
            bind:value={name}
        />
        <br />
        <label for="Description" class="text-xl">Description</label>
        <input
            type="text"
            id="description"
            name="description"
            class="bg-slate-500 w-64"
            bind:value={description}
        />
        <br />
        <a href="javascript:void(0)" class="bg-slate-500 p-2 m-6 rounded-lg text-xl text-center w-64" on:click={createClub}>
            Create Club
        </a>
    </div>
</div>
