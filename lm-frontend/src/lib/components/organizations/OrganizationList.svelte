<script lang="ts">
    import InventoryCellActionButtons from "../inventory/InventoryCellActionButtons.svelte";
    import InventoryList from "../inventory/InventoryList.svelte";

    let headers = [
        {key: "org",        value: "Name"},
        {key: "country",    value: "Country"},
        {key: "contact",    value: "Contact Person"},
        {key: "mail",       value: "Email"},
        {key: "licenses",   value: "Licenses"},
        {key: "actions",    value: "Actions"}
    ]

    let rows = [
        { id: 1, org: "Org", country: "Spain", contact: "Jaime", mail: "jaime.munoz@seqera.io", licenses: 1 },
        { id: 2, org: "SeqeraLabs", country: "Spain", contact: "Boss", mail: "licenses@seqera.io", licenses: 2000 }
    ]

    export let onCreateButton = () => {}

</script>

<InventoryList
    title="Organization Management"
    addButtonText="New Organization"
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