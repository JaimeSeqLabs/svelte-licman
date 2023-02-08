<script lang="ts">
    import LicenseList from "../../lib/components/licenses/LicenseList.svelte";
    import { goto } from "@roxi/routify";
</script>

<LicenseList
    onCreateButton={() => $goto('/licenses/add')}
/>
