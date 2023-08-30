<script lang="ts"> 
    let className: string = "";
    export { className as class };
    
    import { createEventDispatcher } from "svelte";

    const dispatch = createEventDispatcher();
    const handleClick = () => {
        dispatch("click");
    };
</script>

<div
    class={className}
    on:keypress={(e) => e.key == "Enter" && handleClick()}
    on:click={handleClick}
    tabindex="0"
    role="button"
>
    <slot/>
</div>
