<script lang="ts" >


    let msg = "no message"
    function btn(e: Event) {
        console.log(e)
        
        fetch("/api")
            .then(response => response.json())
            .then((data) => {
                console.log(data)
                msg = data.message
            })
    }
  
</script>

<h1>Welcome to SvelteKit</h1>
<button on:click={btn}>Click me</button>

{msg}
