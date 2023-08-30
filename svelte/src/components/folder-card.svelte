<script lang="ts"> 
    export let title: string;
    import { createEventDispatcher } from "svelte";
    import Clickable from "./clickable.svelte";

    const dispatch = createEventDispatcher();
    const handleClick = () => {
        dispatch("click");
    };
</script>

<Clickable
    class="relative border border-transparent hover:border-gray-500 focus:bg-gray-700" 
    on:click={handleClick} 
>
    <img
        class="w-[160px] h-[100px] object-contain"
        src="/Papirus/32x32/places/folder-brown.svg"
        alt={title}
    />
    <div class="p-4 w-[160px]">
        <p
            {title}
            class="block overflow-hidden h-15 mb-2 text-sm tracking-tight break-words text-center"
        >
            {title}
        </p>
    </div>
</Clickable>
