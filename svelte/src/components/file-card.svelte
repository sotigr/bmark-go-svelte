<script lang="ts">
    import { humanDate, humanFileSize } from "$lib"; 
    export let img: string;
    export let title: string;
    export let size: number;
    export let created: Date;

    import { createEventDispatcher } from "svelte";
    import Clickable from "./clickable.svelte";

    const dispatch = createEventDispatcher();
    const handleClick = () => {
        dispatch("click");
    };
</script>

<Clickable
    class="relative border border-transparent hover:border-gray-500 focus:bg-gray-700"
    on:click={() => handleClick}
>
    <img class="w-[160px] h-[100px] object-cover" src={img} alt={title} />
    <div class="p-4 w-[160px] ">
        <p
            {title}
            class="block overflow-hidden h-20 mb-2 text-sm tracking-tight break-words"
        >
            {title}
        </p>
        <div class="flex h-5">
            <div class="flex-1 text-sm">{humanFileSize(size)}</div>
            <div class="flex-1 text-right text-sm">{humanDate(created)}</div>
        </div>
    </div>
</Clickable>
