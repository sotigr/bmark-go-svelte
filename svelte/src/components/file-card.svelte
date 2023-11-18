<script lang="ts">
    import { humanDate, humanFileSize } from "$lib";
    import LazyLoad from "@dimfeld/svelte-lazyload";

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
    $: extMap = {
        png: img, PNG: img, jpg: img,
        jpeg: img, gif: img,  
        
        mp4: "assets/Papirus/24x24/mimetypes/video-x-generic.svg",
        wmv: "assets/Papirus/24x24/mimetypes/video-x-generic.svg",
        flv: "assets/Papirus/24x24/mimetypes/application-vnd.adobe.flash.movie.svg",

        zip: "assets/Papirus/24x24/mimetypes/x-package-repository.svg",
        rar: "assets/Papirus/24x24/mimetypes/x-package-repository.svg",
    };

    $: extension = title.substring(title.lastIndexOf(".") + 1) 
    $: imageName =  (extMap as any)[extension] || "/assets/Papirus/24x24/mimetypes/text-x-generic.svg";
</script>

<Clickable
    class="relative border border-transparent hover:border-gray-500 focus:bg-gray-700"
    on:click={() => handleClick}
>
    <!-- {#if title.endsWith(".png") || title.endsWith(".PNG") || title.endsWith(".jpg") || title.endsWith(".jpeg") || title.endsWith(".gif")} -->
    <LazyLoad height="100px">
        <img
            class="w-[160px] h-[100px] object-contain"
            src={imageName}
            alt={title}
        />
    </LazyLoad>

    <div class="p-4 w-[160px]">
        <p
            {title}
            class="block overflow-hidden h-20 mb-2 text-sm tracking-tight break-words text-center"
        >
            {title}
        </p>
        <div class="flex h-5">
            <div class="flex-1 text-sm">{humanFileSize(size)}</div>
            <div class="flex-1 text-right text-sm">{humanDate(created)}</div>
        </div>
    </div>
</Clickable>
