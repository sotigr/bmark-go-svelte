<script lang="ts">
    import { onMount } from "svelte";
    import {
        Button,
        Breadcrumb,
        BreadcrumbItem,
        Drawer,
        CloseButton,
        ButtonGroup,
        DarkMode,
    } from "flowbite-svelte";
    import { Icon } from "flowbite-svelte-icons";
    import { sineIn } from "svelte/easing";
    import FileCard from "../components/file-card.svelte";
    import FolderCard from "../components/folder-card.svelte";

    let transitionParams = {
        x: -320,
        duration: 200,
        easing: sineIn,
    };

    let events: string[] = [];

    async function handleClick() {
        const response = await fetch("http://localhost/test");
    }

    function arrayGetLast<T>(arr: Array<T>, length: number): Array<T> {
        let arrCopy;
        if (events.length > length) {
            arrCopy = arr.slice(1, arr.length);
        } else {
            arrCopy = [...arr];
        }
        return arrCopy;
    }

    onMount(() => {
        window.eventSource = new EventSource(
            "http://localhost/sys/event-buffer?lobbies=test_lobby:123"
        );
        window.eventSource.addEventListener("test_lobby:test_event", (e) => {
            let eventsCopy = arrayGetLast(events, 30);
            eventsCopy.push(e.data);

            events = eventsCopy; 
        }); 
        window.eventSource.addEventListener("close", (_) => {
            console.log("done");
            window.eventSource.close();
        }); 
    });

    let drawerHidden = true;
</script>

<ButtonGroup>
    <Button on:click={() => (drawerHidden = false)}
        ><Icon name="bars-outline" class="w-5 h-5" /></Button
    >
    <Button>Settings</Button>
    <Button>Messages</Button>
</ButtonGroup>

<Breadcrumb aria-label="Solid background breadcrumb example" solid>
    <BreadcrumbItem href="/" home>Home</BreadcrumbItem>
    <BreadcrumbItem href="/">Projects</BreadcrumbItem>
    <BreadcrumbItem>Flowbite Svelte</BreadcrumbItem>
</Breadcrumb>

<!-- <GradientButton color="pinkToOrange" on:click={handleClick}
    >stream</GradientButton
> -->
<!-- <div>
    {#each events as event}
        <div><p>{event}</p></div>
    {/each}
</div> -->

<div class="grid gap-5 grid-cols-2 md:grid-cols-3 xl:grid-cols-4 2xl:grid-cols-5 py-2">
    <div class="flex justify-center ">
        <FolderCard
            on:click={() => alert("yo")}  
            title="sffs"
        />
    </div>
    <div class="flex justify-center  ">
        <FileCard
            on:click={() => alert("yo")}
            img="/bathuman.png"
            title="1663208150.greenguy212_20220914_190132_1663208150.greenguy212_20220914_190132_1663208150.greenguy212_20220914_190132.jpg "
        />
    </div>
    <div class="flex justify-center ">
        <FileCard
            on:click={() => alert("yo")}
            img="/bathuman.png" 
            title="sffs"
        />
    </div>
    <div class="flex justify-center ">
        <FileCard
            on:click={() => alert("yo")}
            img="/bathuman.png" 
            title="sffs"
        />
    </div>
</div>



<Drawer
    bind:hidden={drawerHidden}
    transitionType="fly"
    {transitionParams}
    id="sidebar1"
>
    <div class="flex items-center">
        <CloseButton
            on:click={() => (drawerHidden = true)}
            class="mb-4 dark:text-white"
        />
    </div>
    <DarkMode />
</Drawer>
