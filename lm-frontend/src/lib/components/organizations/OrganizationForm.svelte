<script lang="ts">
    import { goto } from "@roxi/routify";
    import { FluidForm, Tile, TextInput, ButtonSet, Button } from "carbon-components-svelte";
    import { createOrg } from "../../clients/organizations";

    export let orgName: string = ""
    export let contactName: string = ""
    export let contactMail: string = ""
    export let address: string = ""
    export let postalCode: string = ""
    export let country: string = ""

    let onSubmit = () => {

        createOrg({
            id: "",
            name: orgName,
            contact: contactName,
            mail: contactMail,
            country: country,
            licenses: [],
            date_created: null,
            last_updated: null
        })
        .then(_ => {
            $goto("/organizations")
        })
        .catch(console.error)

    }

</script>


<Tile class="text-xl font-bold">
    Create Organization
</Tile>

<FluidForm>
    <TextInput bind:value={orgName} labelText="Organization name" placeholder="Enter Organization name" required/>
    <TextInput bind:value={contactName} labelText="Contact person" placeholder="Enter Contact person name" required/>
    <TextInput bind:value={contactMail} labelText="Contact email" placeholder="Enter Contact mail" required/>
    <TextInput bind:value={address} labelText="Address" placeholder="Enter Address"/>
    <TextInput bind:value={postalCode} labelText="Postal Code" placeholder="Enter Postal Code"/>
    <TextInput bind:value={country} labelText="Country" placeholder="Enter Country"/> 
</FluidForm>

<br>

<ButtonSet>
    <Button kind="secondary">Cancel</Button>
    <Button kind="primary" on:click={onSubmit}>Save</Button>
</ButtonSet>