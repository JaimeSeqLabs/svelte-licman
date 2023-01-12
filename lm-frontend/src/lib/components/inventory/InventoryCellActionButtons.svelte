<script lang="ts">
    import { Button } from "carbon-components-svelte";
    import type { DataTableCell } from "carbon-components-svelte/types/DataTable/DataTable.svelte";
    import { Archive, Edit, View } from "carbon-icons-svelte";

    export let cell: DataTableCell
    type cellCb = (cell: DataTableCell)=>void

    export let onDetails: cellCb | undefined = undefined
    export let onEdit: cellCb | undefined = undefined
    export let onArchive: cellCb | undefined = undefined
</script>

<Button kind="ghost" icon={View} iconDescription="Details" on:click={(_) => onDetails?.(cell)}/>
<Button kind="ghost" icon={Edit} iconDescription="Edit" on:click={(_) => onEdit?.(cell)}/>
<Button kind="ghost" icon={Archive} iconDescription="Archive" on:click={(_) => onArchive?.(cell)}/>