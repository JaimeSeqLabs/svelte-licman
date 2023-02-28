<script lang="ts">
    import { onMount } from "svelte";
    import { quotas, type Quota } from "../../clients/quota_store";
    import InventoryCellActionButtons from "../inventory/InventoryCellActionButtons.svelte";
    import InventoryList from "../inventory/InventoryList.svelte";
    
    export let onCreateButton = () => {}

    let headers = [
        {key: "id", value: "Name"},
        {key: "default", value: "Default Value"},
        {key: "actions", value: "Actions"}
    ]
    
    let rows: { id:string, default:string }[] = []
    
    quotas.subscribe(qs => {
        rows = qs.map(q => {
            return {
                id: q.name,
                default: q.value
            }
        })
    })

    onMount(() => {
        //fetchQuotas()
    })



</script>

<InventoryList
    title="Quota Management"
    addButtonText="New Quota"
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