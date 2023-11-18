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

    import axios from "axios";
    import type { DirectoryListing } from "../types";
    import Clickable from "../components/clickable.svelte"; 
    let transitionParams = {
        x: -320,
        duration: 200,
        easing: sineIn,
    };

    // let events: string[] = [];

    // async function handleClick() {
    //     const response = await fetch("http://localhost/test");
    // }

    // function arrayGetLast<T>(arr: Array<T>, length: number): Array<T> {
    //     let arrCopy;
    //     if (events.length > length) {
    //         arrCopy = arr.slice(1, arr.length);
    //     } else {
    //         arrCopy = [...arr];
    //     }
    //     return arrCopy;
    // }

    let listing: DirectoryListing | null = null;
    let cPath: string = "";
 

    async function listDirectory() {
        let path = sessionStorage.getItem("path") || "";
        const res = await axios.get(`http://localhost/list?path=${path}`);
        listing = res.data;
        cPath = path;
        window.scrollTo(0,0)  
    }

    async function setDirectory(path: string) {
        sessionStorage.setItem("path", path);
        await listDirectory();
        cPath = path;
    }

    function goBack() {
        const clean = cPath.substring(0, cPath.length-1)
            const sub = clean.substring(0, clean.lastIndexOf("/")) 
            if (sub == "") {
                setDirectory("")
            } else {
                setDirectory(sub+"/")
            }
    }
 
    onMount(async () => {
        // window.eventSource = new EventSource(
        //     "http://localhost/sys/event-buffer?lobbies=test_lobby:123"
        // );
        // window.eventSource.addEventListener("test_lobby:test_event", (e) => {
        //     let eventsCopy = arrayGetLast(events, 30);
        //     eventsCopy.push(e.data);

        //     events = eventsCopy;
        // });
        // window.eventSource.addEventListener("close", (_) => {
        //     console.log("done");
        //     window.eventSource.close();
        // });
        await listDirectory();
    });

    let drawerHidden = true;

    $: pathParts = cPath.split("/");
</script>
<div class="pb-[40px]"> 
    <div class="fixed h-[40px] w-full bottom-0 z-10 bg-black flex items-center justify-end px-2 lg:hidden">
        <Clickable on:click={goBack}>
            <Icon name="angle-left-outline" class="w-5 h-5" />
        </Clickable>
    </div>
    <ButtonGroup>
        <Button on:click={() => (drawerHidden = false)}
            ><Icon name="bars-outline" class="w-5 h-5" /></Button
        >
        <!-- <Button>Settings</Button>
        <Button>Messages</Button> -->
    </ButtonGroup>
    
    <Breadcrumb
        aria-label="Solid background breadcrumb example"
        olClass="flex flex-wrap gap-2 p-2 border border-gray-500 rounded-xl w-full"
    >
        <Clickable class="pt-1" on:click={() => setDirectory("")}>
            <BreadcrumbItem home>Home</BreadcrumbItem>
        </Clickable>
        {#if pathParts.length > 1}
            {#each pathParts as pathPart, i}
                <Clickable
                    class="pt-1"
                    on:click={() => {
                        setDirectory(pathParts.slice(0, i + 1).join("/") + "/");
                    }}
                >
                    <BreadcrumbItem>{pathPart}</BreadcrumbItem>
                </Clickable>
            {/each}
        {/if}
    </Breadcrumb>
    
    <!-- <GradientButton color="pinkToOrange" on:click={handleClick}
        >stream</GradientButton
    > -->
    <!-- <div>
        {#each events as event}
            <div><p>{event}</p></div>
        {/each}
    </div> -->
    
    <div class="flex gap-2 overflow-hidden flex-wrap p-2">
        {#if listing}
            {#each listing.folders as item}
                <div class="">
                    <FolderCard
                        on:click={() => setDirectory(item.fullName)}
                        title={item.name}
                    />
                </div>
            {/each}
            {#each listing.files as item}
                <div class="flex justify-center">
                    <div class="flex justify-center">
                        <FileCard
                            on:click={() => alert("yo")}
                            size={item.size}
                            created={item.created}
                            img={`http://localhost:80/read?path=${item.fullName}`}
                            title={item.name}
                        />
                    </div>
                </div>
            {/each}
        {/if}
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
</div>
