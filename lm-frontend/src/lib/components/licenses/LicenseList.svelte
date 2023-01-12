<script lang="ts">
    import { 
        StructuredList,
        StructuredListHead,
        StructuredListBody,
        StructuredListRow,
        StructuredListCell
    } from "carbon-components-svelte";

    let now = new Date().toUTCString()
    let licenseList = [
        { ID: "id1", creationDate: now, org: "Org1" },
        { ID: "id2", creationDate: now, org: "Org2" },
        { ID: "id3", creationDate: now, org: "Org3" },
    ]

</script>

<StructuredList>

    <StructuredListHead>
        <StructuredListRow head>

            <StructuredListCell head>ID</StructuredListCell>
            <StructuredListCell head>Created at</StructuredListCell>
            <StructuredListCell head>Organization</StructuredListCell>

        </StructuredListRow>
    </StructuredListHead>


    <StructuredListBody>

        {#each licenseList as lic}
        <StructuredListRow>
            <StructuredListCell>{lic.ID}</StructuredListCell>
            <StructuredListCell>{lic.creationDate}</StructuredListCell>
            <StructuredListCell>{lic.org}</StructuredListCell>
        </StructuredListRow>
        {/each}    
    
    </StructuredListBody>


    
</StructuredList>