<script lang="ts">
    import { goto } from "@roxi/routify";
    import { FluidForm, Tile, TextInput, ButtonSet, Button } from "carbon-components-svelte";
    import { createNewProduct } from "../../clients/product";

    export let skuCode: string = ""
    export let productName: string = ""
    export let productDescription: string = ""
    export let productInstructions: string = ""

    let onSubmit = () => {
        createNewProduct({
            sku: skuCode,
            name: productName,
            install_instructions: productInstructions
        })
        .then(_ => {
            $goto("/products")
        })
        .catch(console.error)
    }

</script>


<Tile class="text-xl font-bold">
    Create Product
</Tile>

<FluidForm>
    <TextInput bind:value={skuCode} labelText="SKU code" required/>
    <TextInput bind:value={productName} labelText="Product name" required/>
    <TextInput bind:value={productDescription} labelText="Product description"/>
    <TextInput bind:value={productInstructions} labelText="Product install instructions"/>
</FluidForm>

<br>

<ButtonSet>
    <Button kind="secondary">Cancel</Button>
    <Button kind="primary" on:click={onSubmit}>Save</Button>
</ButtonSet>