<script lang="ts">
    import InventoryCellActionButtons from "../inventory/InventoryCellActionButtons.svelte";
    import InventoryList from "../inventory/InventoryList.svelte";

    let headers = [
        {key: "id",         value: "SKU"},
        {key: "product",    value: "Name"},
        {key: "actions",    value: "Actions"}
    ]

    let rows = [
        { id: "TWR", product: "Tower" },
        { id: "XPACX-AMZN", product: "AWS xpack-plugin" },
        { id: "XPACX-AZRE", product: "Azure xpack-plugin" }
    ]

    export let onCreateButton = () => {}

</script>

<InventoryList
    title="Product Management"
    addButtonText="New Product"
    on:new={onCreateButton}
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