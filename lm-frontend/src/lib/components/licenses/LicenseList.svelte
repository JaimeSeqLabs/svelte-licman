<script lang="ts">

import InventoryList from "../inventory/InventoryList.svelte";
    import InventoryCellActiveTag from "../inventory/InventoryCellActiveTag.svelte";
    import InventoryCellActionButtons from "../inventory/InventoryCellActionButtons.svelte";
    import { listAllLicenses, type DomainLicense, type ListAllLicensesItem } from "../../clients/license";
    import { onMount } from "svelte";
    import { describeOrg, listAllOrgs, type DescribeOrgResponse } from "../../clients/organizations";
    import { Button, Modal, Tile } from "carbon-components-svelte";
    import { LM_PRIVATE_API } from "../../clients/common";


    export let onCreateButton = () => {}
    let showLicenseDetailsModal = false
    let licenseDetailLicense:DomainLicense

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

    function showDetailsFor(id:string) {
        
        licenseDetailLicense = licenses.find(lic => lic.id == id)
        console.log(`show details for ${id}, ${licenseDetailLicense.id}`);
        showLicenseDetailsModal = true
    }
    
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
    let:row
>
    <!-- Display license status with a tag -->
    {#if cell.key == "active"}
        <InventoryCellActiveTag {cell}/>
    <!-- Display action buttons-->
    {:else if cell.key == "actions"}
        <InventoryCellActionButtons
            {cell} 
            onDetails={()=>showDetailsFor(row["id"])}
        />
    <!-- Display default cell value-->
    {:else}
        {cell.value}
    {/if}

</InventoryList>

<Modal
    passiveModal
    bind:open={showLicenseDetailsModal}
    modalHeading="License Details"
    on:close={()=>showLicenseDetailsModal=false}
>

<Tile> ID: {licenseDetailLicense?.id}</Tile>
<Tile> Status: {licenseDetailLicense?.status}</Tile>
<Tile> Organization: {licenseDetailLicense?.organization_id}</Tile>
<Tile> Activation: {licenseDetailLicense?.activation_date}</Tile>
<Tile> Expiration: {licenseDetailLicense?.expiration_date}</Tile>
<Tile> Contact: {licenseDetailLicense?.activation_date}</Tile>

<Button on:click={()=>window.open(LM_PRIVATE_API + `/licenses/download/${licenseDetailLicense?.id}`)}>Download</Button>

</Modal>