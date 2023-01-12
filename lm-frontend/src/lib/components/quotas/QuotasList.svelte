<script lang="ts">
    import InventoryCellActionButtons from "../inventory/InventoryCellActionButtons.svelte";
    import InventoryList from "../inventory/InventoryList.svelte";

    let now = new Date().toUTCString()

    let headers = [
        {key: "id", value: "Name"},
        {key: "default", value: "Default Value"},
        {key: "lastUpdate", value: "Last Update"},
        {key: "actions", value: "Actions"}
    ]

    let rows = [
        { id: "maxUsersPerWsp", default: 10, lastUpdate: now },
        { id: "maxWspPerOrg", default: 15, lastUpdate: now }
    ]

</script>

<InventoryList
    title="Quota Management"
    addButtonText="New Quota"
    {headers}
    {rows}
    let:cell
>
    <!-- Display action buttons-->
    {#if cell.key == "actions"}
        <InventoryCellActionButtons {cell}/>
    <!-- Display default cell value-->
    {:else}
        {cell.value}
    {/if}

</InventoryList>