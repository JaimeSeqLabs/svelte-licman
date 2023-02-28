<script lang="ts">
    import { onMount } from "svelte";
    import { describeOrg, listAllOrgs, type DescribeOrgResponse } from "../../clients/organizations";
    import InventoryCellActionButtons from "../inventory/InventoryCellActionButtons.svelte";
    import InventoryList from "../inventory/InventoryList.svelte";

    export let onCreateButton = () => {}

    let headers = [
        {key: "org",        value: "Name"},
        {key: "country",    value: "Country"},
        {key: "contact",    value: "Contact Person"},
        {key: "mail",       value: "Email"},
        {key: "licenses",   value: "Licenses"},
        {key: "actions",    value: "Actions"}
    ]

    let orgs: DescribeOrgResponse[] = []
    function fetchOrgs() {
        listAllOrgs()
        .then(async (list) => {

            orgs = (
                await Promise.all(
                    list.data.organizations.map(org => describeOrg(org.id))
                )
            ).map(desc => desc.data)
            
        })
        .catch(err => {
            console.error(err)
        })
    }

    let rows: { id:string, org:string, country:string, contact:string, mail: string, licenses:number }[] = []
    $: rows = orgs.map(org => {
        return {
            id: org.id,
            org: org.name,
            country: org.country,
            contact: org.contact,
            mail: org.mail,
            licenses: org.licenses.length
        }
    })

    onMount(() => {
        fetchOrgs()
    })

  

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