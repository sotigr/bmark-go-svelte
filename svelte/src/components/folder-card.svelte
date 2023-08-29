<script lang="ts">
    import { Card } from "flowbite-svelte";   
    export let title: string;
    import { createEventDispatcher } from "svelte";

    const dispatch = createEventDispatcher();
    const handleClick = () => {
        dispatch("click");
    };
</script>

<div class="relative">
    
    <Card href="#" class="pb-5"  on:click={handleClick} padding="none">
        <img class="w-[185px] h-[128px] object-contain" src="/Papirus/32x32/places/folder-brown.svg" alt={title}/>
        <div class="p-4 w-[185px]">
            <p
                {title}
                class="block overflow-hidden h-20 mb-2 text-lg font-bold tracking-tight break-words"
            >
                {title}
            </p> 
        </div>
    </Card>
</div>
