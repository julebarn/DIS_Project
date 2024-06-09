<script lang="ts">
    import { register } from "$lib/auth";

    let username = "";
    let password = "";

    function onRegister() {
        // Validate that username and password are nonempty
        if (!username || !password) {
            alert('Please enter a username and password');
            // TODO: show error message
            return;
        }

        // Validate that username is alphanumeric
        let validated = true; 
        if (/^([A-Za-z0-9]){5,}$/.test(username) === false) {
            alert('Invalid Username: Username must be at least 5 characters long and contain only letters and numbers');
            validated = false;
        }

        // Instead of matching if a regex is true, we will match if it is false instead,
        // Cases: 
        // [^A-Z]* matches any string that does NOT contain an uppercase letter
        // [^a-z]* matches any string that does NOT contain a lowercase letter
        // [^0-9]* matches any string that does NOT contain a number
        // [^!@#$%^&*(),./+-]* matches any string that does NOT contain a 'special' character, as defined by the regex
        // .{0,7} matches any string that is LESS THAN 8 characters long
        // We OR each case together, so if one (or more) of the cases is true, then the password is invalid

        if (/^([^A-Z]*|[^a-z]*|[^0-9]*|[^!@#$%^&*(),./+-]*|.{0,7})$/.test(password) === true) {
            alert('Invalid Password: Password must be at least 8 characters long and contain at least one uppercase letter, one lowercase letter, one number, and one special character');
            validated = false;
        }
        
        if (validated === false) {
            return;
        }

        register(username, password);
    }
</script>

<div class="flex flex-col items-stretch bg-slate-400 w-screen min-h-screen">
    <h1 class="text-5xl text-center m-5">Register</h1>

    <div class="flex flex-col justify-center items-center space-y-1">
        <label for="username" class ="text-xl">Username</label>
        <input class="bg-slate-500 w-64" type="text" id="username" name="username" bind:value={username} />
        <label for="password" class ="text-xl">Password</label>
        <input class="bg-slate-500 w-64" type="password" id="password" name="password" bind:value={password} />

        <!-- Kinda janky but it works! -->
        <a href="javascript:void(0)" class="bg-slate-500 p-2 m-6 rounded-lg text-xl text-center w-64" on:click={onRegister}>
                Register
        </a>
    </div>
</div>
