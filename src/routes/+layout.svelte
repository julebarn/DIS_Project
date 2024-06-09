<script>
    import "../app.css";

    import { refreshTokens } from "$lib/auth";
    import { onMount } from "svelte";

    onMount(() => {
        refreshTokens();
    });
</script>

<slot />
