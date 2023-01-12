<script lang="ts">
    import type { DataTableCell } from "carbon-components-svelte/types/DataTable/DataTable.svelte";
    import { Tag } from "carbon-components-svelte";

    export let cell: DataTableCell
</script>

{#if cell.value == true}
    <Tag type="green">ACTIVE</Tag>
{:else}
    <Tag type="red">SUSPENDED</Tag>
{/if}
