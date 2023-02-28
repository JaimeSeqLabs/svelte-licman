<script lang="ts">
    import {
        DataTable,
        Toolbar,
        ToolbarContent,
        ToolbarSearch,
        Toggle,
        Button,
        Pagination,
        Tag
    } from "carbon-components-svelte";
    import { AddAlt } from "carbon-icons-svelte";
    import type { DataTableHeader, DataTableRow } from "carbon-components-svelte/types/DataTable/DataTable.svelte";
    import { createEventDispatcher } from "svelte";

    export let title: string
    export let addButtonText: string = "New"

    export let headers: DataTableHeader[]
    export let rows: DataTableRow[]

    export let page = 1
    export let pageSize = 10
    export let pageSizes: number[] = [5, 10, 15]

    const dispatch = createEventDispatcher()

    let dispatchCreateEvent = () => dispatch("new")


</script>

<DataTable
    {title}
    stickyHeader
    size="compact"
    bind:headers
    bind:rows
>

    <Toolbar>
        <ToolbarContent>
            <ToolbarSearch class="mr-3" persistent/>
            <Toggle class="mr-32" size="sm" labelA="Show archived" labelB="Show archived"/>
            <Button icon={AddAlt} on:click={dispatchCreateEvent}>{addButtonText}</Button>
        </ToolbarContent>
    </Toolbar>

    <!-- pass cell down to new slot -->
    <svelte:fragment slot="cell" let:cell let:row>

        <slot cell={cell} row={row}/>

    </svelte:fragment>

</DataTable>

<Pagination
    bind:pageSize
    bind:page
    {pageSizes}
/>