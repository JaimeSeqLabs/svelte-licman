<script lang="ts">
    import { onMount } from "svelte";
    import { listAllProducts, type DescribeProductResponse } from "../../clients/product";
    import InventoryCellActionButtons from "../inventory/InventoryCellActionButtons.svelte";
    import InventoryList from "../inventory/InventoryList.svelte";
    
    export let onCreateButton = () => {}

    let headers = [
        {key: "id",         value: "SKU"},
        {key: "product",    value: "Name"},
        {key: "actions",    value: "Actions"}
    ]

    let products: DescribeProductResponse[] = []
    function fetchProducts() {
        listAllProducts()
        .then(res => products = res.data.products)
        .catch(console.error)
    }

    let rows: { id:string, product:string }[] = []
    $: rows = products.map(prod => {
        return {
            id: prod.sku,
            product: prod.name
        }
    })

    onMount(() => {
        fetchProducts()
    })


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