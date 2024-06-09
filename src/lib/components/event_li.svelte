<script lang="ts">
export let name = 'world';
export let start_time = '00:00';
export let ID = '0';
</script>

<a href="/event?id={ID}">
<div class="bg-slate-500 p-9 m-2 rounded-lg grid grid-cols-2  text-xl">
    <h2 class="place-self-start">{name}</h2>
    <p class="place-self-end">{start_time}</p>
</div>
</a>