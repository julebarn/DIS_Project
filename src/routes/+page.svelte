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

<h1 class="text-8xl">Welcome to dis</h1>
<button on:click={btn}>Click me</button>

{msg}
