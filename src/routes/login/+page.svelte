<script lang="ts">
  import { login } from '$lib/auth';

  let username = '';
  let password = '';

  async function handleSubmit() {
    // Validate that username and password are nonempty
    if (!username || !password) {
      alert('Please enter a username and password');
      // TODO: show error message
      return;
    }
    
    login(username, password);
  }
</script>

<div class="flex flex-col items-stretch bg-slate-400 w-screen min-h-screen">
  <h1 class="text-5xl text-center m-5">Login</h1>

  <div class="flex flex-col justify-center items-center space-y-1">
      <label for="username" class ="text-xl">Username</label>
      <input class="bg-slate-500 w-64" type="text" id="username" name="username" bind:value={username} />
      <label for="password" class ="text-xl">Password</label>
      <input class="bg-slate-500 w-64" type="password" id="password" name="password" bind:value={password} />

      <!-- Kinda janky but it works! -->
      <a href="javascript:void(0)" class="bg-slate-500 p-2 m-6 rounded-lg text-xl text-center w-64" on:click={handleSubmit}>
              Login
      </a>
  </div>
</div>