<script lang="ts">

import InventoryList from "../inventory/InventoryList.svelte";
    import InventoryCellActiveTag from "../inventory/InventoryCellActiveTag.svelte";
    import InventoryCellActionButtons from "../inventory/InventoryCellActionButtons.svelte";
    import { listAllLicenses, type DomainLicense, type ListAllLicensesItem } from "../../clients/license";
    import { onMount } from "svelte";
    import { describeOrg, listAllOrgs, type DescribeOrgResponse } from "../../clients/organizations";


    export let onCreateButton = () => {}

    let headers = [
        {key: "id",         value: "ID"},
        {key: "org",        value: "Company"},
        {key: "activation", value: "Activation Date"},
        {key: "expiration", value: "Expiration Date"},
        {key: "active",     value: "Status"},
        {key: "actions",    value: "Actions"}
    ]

    let licenses: DomainLicense[] = []
    function fetchLicenses() {
        listAllLicenses()
        .then(res => {
            licenses = res.licenses
        })
        .catch(console.error)
    }

    let orgs: DescribeOrgResponse[] = []
    function fetchOrgs() {
        listAllOrgs()
        .then(async (list) => {
            orgs = (
                await Promise.all(
                    list.data.organizations.map(org => describeOrg(org.id))
                )
            )
            .map(desc => desc.data)
        })
        .catch(err => {
            console.error(err)
        })
    }

    let rows: { id:string, org:string, activation:string, expiration:string, active:boolean}[] = []
    $: rows = licenses.map(lic => {        
        return {
            id: lic.id,
            org: orgs.find(org => org.id == lic.organization_id)?.name, // TODO: licenses are ready before orgs, orgs may be empty
            activation: lic.activation_date,
            expiration: lic.expiration_date,
            active: lic.status == "active"
        }
    })
    
    onMount(() => {
        fetchLicenses()
        fetchOrgs()
    })


</script>

<InventoryList
    title="License Manager"
    addButtonText="New License"
    on:new={onCreateButton}
    {headers}
    {rows}
    let:cell
>
    <!-- Display license status with a tag -->
    {#if cell.key == "active"}
        <InventoryCellActiveTag {cell}/>
    <!-- Display action buttons-->
    {:else if cell.key == "actions"}
        <InventoryCellActionButtons {cell}/>
    <!-- Display default cell value-->
    {:else}
        {cell.value}
    {/if}

</InventoryList>