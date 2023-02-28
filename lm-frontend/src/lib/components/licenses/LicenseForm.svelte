<script lang="ts">
    import { goto } from "@roxi/routify";
    import { Tile, TextInput, ButtonSet, Button, Select, SelectItem, Toggle, Modal, DatePicker, DatePickerInput, TextArea, TextAreaSkeleton, DataTable, NumberInput } from "carbon-components-svelte";
    import { TrashCan } from "carbon-icons-svelte";
    import { onMount } from "svelte";
    import { createNewLicense } from "../../clients/license";
    import { describeOrg, listAllOrgs, type DescribeOrgResponse, type ListAllOrgsItem } from "../../clients/organizations";
    import { listAllProducts, type DescribeProductResponse } from "../../clients/product";
    import { quotas, type Quota } from "../../clients/quota_store";

    // confirm suspend license modal
    let licenseIsSuspended = false
    let showSuspendLicenseModal = false
    let toggle = false
    let handleSuspendLicense = (toggled: boolean) => {
        if (toggled) {
            showSuspendLicenseModal = true
        }
    }

    let prods: DescribeProductResponse[] = []
    function fetchProds() {
        listAllProducts()
        .then(res => {
            prods = res.data.products
            // add empty to trigger select
            prods.push({id: "",sku: "",name: "",install_instructions: "",
                        license_count: 0,date_created: "",last_updated: ""})
        })
        .catch(err => {
            console.error(err)
        })
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
            
            // add empty to trigger select
            orgs.push({id: "", name: "",contact: "",mail: "",country: "",
                    licenses: [],date_created: "",last_updated: ""})
            
        })
        .catch(err => {
            console.error(err)
        })
    }

    // component state
    export let prodName = ""
    export let orgName = ""
    export let contactName = ""
    export let contactMail = ""
    export let licenseNotes = ""
    export let activationDate = ""
    export let expirationDate = ""
    export let availableQuotas:Quota[] = []
    export let customQuotas:Quota[] = []
    let selectedCustomQuotaname:string = ""

    $: {
        let current = orgs.find(org => org.name == orgName)
        if (current) {
            contactName = current.contact
            contactMail = current.mail
        }
    }

    quotas.subscribe(qs => {
        availableQuotas = [{name:"", value:""}, ...qs]
    })

    // custom quotas table
    let headers = [
        {key: "name", value: "Name"},
        {key: "value", value: "Value"},
        {key: "actions",    value: "Actions"}
    ]

    // add selected quota to custom list
    $: if (selectedCustomQuotaname != "") {
        if (!customQuotas.find(q=>q.name == selectedCustomQuotaname)) {
            let q = availableQuotas.find(quota => quota.name == selectedCustomQuotaname)
            customQuotas = [...customQuotas, q]
        }
        selectedCustomQuotaname = ""
    }

    // update table rows when customQuotas change
    let rows:{id: string,name: string,value:string}[] = []
    $: {
        rows = customQuotas.map(q => {
            return {
                id: q.name,
                name: q.name,
                value: q.value
            }
        })
    }

    function deleteCustomQuota(name: string) {
        customQuotas = customQuotas.filter(q => q.name != name)
    }

    function onSubmit() {
        createNewLicense({
            features: "",
            status: toggle ? "suspended" : "active",
            version: "",
            note: licenseNotes,
            contact: contactName,
            mail: contactMail,
            product_skus: [ prods.find(p => p.name == prodName).sku ],
            organization_name: orgName,
            quotas: customQuotas.reduce((acc, q)=> {
                acc[q.name]=q.value
                return acc
            }, {}),
            expiration_date: expirationDate,
            activation_date: activationDate
        })
        .then(_ => {
            $goto("/licenses")
        })

    }
    
    onMount(async () => {
        fetchProds()
        fetchOrgs()
    })

</script>


<Tile class="text-xl font-bold">
    Create License
</Tile>

<Tile>

    <Select labelText="Product" bind:selected={prodName}>
        {#each prods as prod}
            <SelectItem value={prod.name}/>
        {/each}
    </Select>

    <Select labelText="Company" bind:selected={orgName}>
        {#each orgs as org}
            <SelectItem value={org.name}/>
        {/each}
    </Select>

    <TextInput bind:value={contactName} labelText="Contact person" readonly/>
    <TextInput bind:value={contactMail} labelText="Contact email" readonly/>

    <Toggle 
        labelText="Suspended license"
        bind:toggled={toggle}
        on:toggle={evt => handleSuspendLicense(evt.detail.toggled)}
    />
    <br>

    <DatePicker 
        datePickerType="range" 
        dateFormat="d/m/Y"
        bind:valueFrom={activationDate}
        bind:valueTo={expirationDate}
    >
        <DatePickerInput labelText="Activation date" placeholder="dd/mm/yyyy" />
        <DatePickerInput labelText="Expiration date" placeholder="dd/mm/yyyy" />
    </DatePicker>

    <br>

    <TextArea labelText="License Notes" bind:value={licenseNotes}/>

    
    <!-- Custom quotas-->

    <Select
        labelText="Select custom quotas"
        bind:selected={selectedCustomQuotaname}
    >
        {#each availableQuotas as q}
            <SelectItem value={q.name}/>
        {/each}
    </Select>

    <DataTable
        {headers}
        {rows}
    >
        <svelte:fragment slot="cell" let:cell let:row>

            {#if cell.key == "actions"}
                <Button kind="ghost" icon={TrashCan} iconDescription="Delete" on:click={() => deleteCustomQuota(row["name"])}/>
            
            {:else if cell.key == "value"}
                <NumberInput value={cell.value}/>
            
            {:else}
                {cell.value}
            
            {/if}

        </svelte:fragment>

    </DataTable>


</Tile>

<br>

<ButtonSet>
    <Button kind="secondary">Cancel</Button>
    <Button kind="primary" on:click={onSubmit}>Save</Button>
</ButtonSet>

<Modal
    primaryButtonText="OK"
    on:click:button--primary={()=>{
        licenseIsSuspended = true
        showSuspendLicenseModal = false
    }}
    secondaryButtonText="Cancel"
    on:click:button--secondary={()=>{
        licenseIsSuspended = false
        showSuspendLicenseModal = false
        toggle = false
    }}
    bind:open={showSuspendLicenseModal}
>
    <p>Are you sure you want to suspend this license?</p>
</Modal>