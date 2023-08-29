<script lang="ts">
    import { Card } from "flowbite-svelte";
    import { SpeedDial, Listgroup, ListgroupItem } from "flowbite-svelte";
    import { Icon } from "flowbite-svelte-icons";
    export let img: string; 
    export let title: string;
    import { createEventDispatcher } from "svelte";

    const dispatch = createEventDispatcher();
    const handleClick = () => {
        dispatch("click");
    };
</script>

<div class="relative">
    <SpeedDial
        defaultClass="absolute right-6 top-6"
        tooltip="none"
        pill={true}
        placement="top-end"
        gradient={true}
    >
        <Icon name="dots-horizontal-outline" slot="icon" class="w-8 h-8" />
        <Listgroup class="divide-none" active>
            <ListgroupItem class="flex gap-2 md:px-5">
                <Icon name="download-solid" class="w-3.5 h-3.5" />
                Download
            </ListgroupItem>
            <ListgroupItem class="flex gap-2 md:px-5">
                <Icon name="pen-solid" class="w-3.5 h-3.5" />
                Rename
            </ListgroupItem>
            <ListgroupItem class="flex gap-2 md:px-5">
                <Icon name="file-paste-solid" class="w-3.5 h-3.5" />
                Move
            </ListgroupItem>
            <ListgroupItem class="flex gap-2 md:px-5">
                <Icon name="trash-bin-solid" class="w-3.5 h-3.5" />
                Delete
            </ListgroupItem>
        </Listgroup>
    </SpeedDial>
    <Card href="#" on:click={handleClick} padding="none"   >
        <img class="w-[185px] h-[128px] object-cover" src={img} alt={title}/>
        <div class="p-4 w-[185px]">
            <p
                {title}
                class="block overflow-hidden h-20 mb-2 text-sm font-bold tracking-tight break-words"
            >
                {title}
            </p>
            <div class="flex h-5">
                <div class="flex-1">100 Mb</div>
                <div class="flex-1 text-right">19/9/2022</div>
            </div>
        </div>
    </Card>
</div>
