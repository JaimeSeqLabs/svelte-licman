<script lang="ts">
    import { goto } from "@roxi/routify";
    import { Tile, TextInput, ButtonSet, Button, Select, SelectItem, Toggle, Modal, DatePicker, DatePickerInput, TextArea, TextAreaSkeleton } from "carbon-components-svelte";
    import { onMount } from "svelte";
    import { createNewLicense } from "../../clients/license";
    import { describeOrg, listAllOrgs, type DescribeOrgResponse, type ListAllOrgsItem } from "../../clients/organizations";
    import { listAllProducts, type DescribeProductResponse } from "../../clients/product";


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

    $: {
        let current = orgs.find(org => org.name == orgName)
        if (current) {
            contactName = current.contact
            contactMail = current.mail
        }
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
            quotas: {},
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

    <!-- TODO: Custom quotas-->


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