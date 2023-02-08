<script lang="ts">
    import InventoryList from "../inventory/InventoryList.svelte";
    import InventoryCellActiveTag from "../inventory/InventoryCellActiveTag.svelte";
    import InventoryCellActionButtons from "../inventory/InventoryCellActionButtons.svelte";

    let now = new Date().toUTCString()

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

    export let onCreateButton = () => {}

</script>

<InventoryList
    title="License Manager"
    addButtonText="New License"
    on:new={onCreateButton}
    {headers}
    {rows}
    let:cell
>
    <!-- Display license status with a tag -->
    {#if cell.key == "active"}
        <InventoryCellActiveTag {cell}/>
    <!-- Display action buttons-->
    {:else if cell.key == "actions"}
        <InventoryCellActionButtons {cell}/>
    <!-- Display default cell value-->
    {:else}
        {cell.value}
    {/if}

</InventoryList>