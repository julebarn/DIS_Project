<script lang="ts">
    import { goto } from '$app/navigation';
    import { logout } from '$lib/auth';
    
    logout();
    goto('/');
</script>