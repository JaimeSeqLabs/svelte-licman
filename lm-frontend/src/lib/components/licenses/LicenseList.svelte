<script lang="ts">
    import { 
        AddAlt,
        ViewFilled,
        Edit,
        Download,
        View
    } from "carbon-icons-svelte";
    import {
        DataTable,
        Toolbar,
        ToolbarContent,
        ToolbarSearch,
        Toggle,
        Button,
        ButtonSet,
        Pagination,
        Tag
    } from "carbon-components-svelte";

    let now = new Date().toUTCString()

    let licenseList = [
        { ID: "id1", creationDate: now, org: "Org1" },
        { ID: "id2", creationDate: now, org: "Org2" },
        { ID: "id3", creationDate: now, org: "Org3" },
    ]

    let headers = [
        {key: "id",         value: "ID"},
        {key: "org",        value: "Company"},
        {key: "activation", value: "Activation Date"},
        {key: "expiration", value: "Expiration Date"},
        {key: "active",     value: "Status"},
        {key: "actions",    value: "Actions"}
    ]

    let rows = [
        { id: "001", org: "Evil Corp", activation: now, expiration: now, active: true },
        { id: "002", org: "Seqera Labs", activation: now, expiration: now, active: false }
    ]

    let pageSize = 5
    let page = 1

</script>


<DataTable
    title="License Manager"
    stickyHeader
    size="compact"
    bind:headers
    bind:rows
>

    <Toolbar>
        <ToolbarContent>
            <ToolbarSearch persistent/>
            <Toggle class="mr-32" size="sm" labelA="Show archived" labelB="Show archived"/>
            <Button icon={AddAlt}>Add new</Button>
        </ToolbarContent>
    </Toolbar>

    <svelte:fragment slot="cell" let:cell>

        <!-- Display license status with a tag -->
        {#if cell.key == "active"}
            {#if cell.value == true}
                <Tag type="green">ACTIVE</Tag>
            {:else}
                <Tag type="red">SUSPENDED</Tag>
            {/if}
        
        <!-- Display action buttons-->
        {:else if cell.key == "actions"}
            <Button kind="ghost" icon={View} iconDescription="View"/>
            <Button kind="ghost" icon={Edit} iconDescription="Edit"/>
            <Button kind="ghost" icon={Download} iconDescription="Download"/>
        
        <!-- Display default cell value-->
        {:else}
            {cell.value}
        {/if}

    </svelte:fragment>

</DataTable>

<Pagination
    bind:pageSize
    bind:page
    pageSizes={[5, 10, 15]}
/>
