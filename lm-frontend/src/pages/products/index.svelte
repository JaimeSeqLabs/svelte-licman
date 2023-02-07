<script lang="ts">
    import ProductList from "../../lib/components/products/ProductList.svelte";
    import { goto } from "@roxi/routify";

</script>

<ProductList onCreateButton={() => $goto('/products/add')}/>