<script lang="ts">
    import OrganizationList from "../../lib/components/organizations/OrganizationList.svelte";
    import { goto } from "@roxi/routify";

</script>

<OrganizationList
    onCreateButton={() => $goto('/organizations/add')}
/>