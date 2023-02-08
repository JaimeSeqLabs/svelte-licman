<script lang="ts">
    import { Tile, TextInput, ButtonSet, Button, Select, SelectItem, Toggle, Modal, DatePicker, DatePickerInput, TextArea, TextAreaSkeleton } from "carbon-components-svelte";

    const availableProducts = [
        "",
        "TWR",
        "XPACK"
    ]

    const orgData = [
        { name: "", person: "", mail: "" },
        { name: "Org1", person: "Alice", mail: "alice@mail.com" },
        { name: "Org2", person: "Bob", mail: "bob@mail.com" },
    ]
    
    export let skuCode = ""
    export let orgName = ""
    
    let currentOrg = { name: "", person: "", mail: "" }
    $: currentOrg = orgData.find(org => org.name == orgName)

    let licenseIsSuspended = false
    let showSuspendLicenseModal = false
    let toggle = false
    let handleSuspendLicense = (toggled: boolean) => {
        if (toggled) {
            showSuspendLicenseModal = true
        }
    }

    let onSubmit = () => {
        console.log({
        });
    }

</script>


<Tile class="text-xl font-bold">
    Create License
</Tile>

<Tile>

    <Select labelText="Product" bind:selected={skuCode}>
        {#each availableProducts as prod}
            <SelectItem value={prod}/>
        {/each}
    </Select>

    <Select labelText="Company" bind:selected={orgName}>
        {#each orgData as org}
            <SelectItem value={org.name}/>
        {/each}
    </Select>

    <TextInput bind:value={currentOrg.person} labelText="Contact person" readonly />
    <TextInput bind:value={currentOrg.mail} labelText="Contact email" readonly />

    <Toggle 
        labelText="Suspended license"
        bind:toggled={toggle}
        on:toggle={evt => handleSuspendLicense(evt.detail.toggled)}
    />
    <br>

    <DatePicker datePickerType="range">
        <DatePickerInput labelText="Activation date" placeholder="mm/dd/yyyy" />
        <DatePickerInput labelText="Expiration date" placeholder="mm/dd/yyyy" />
    </DatePicker>

    <br>

    <TextArea labelText="License Notes"/>

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