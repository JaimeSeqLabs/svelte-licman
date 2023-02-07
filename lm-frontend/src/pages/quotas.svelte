<script lang="ts">
    import { Modal, NumberInput, TextInput, TextInputSkeleton } from "carbon-components-svelte";
    import QuotasList from "../lib/components/quotas/QuotasList.svelte";

    let openModal = false

</script>

<QuotasList onCreateButton={()=>{ openModal = true }}/>

<Modal
    bind:open={openModal}
    modalHeading="Create Quota"
    primaryButtonText="Save"
    secondaryButtonText="Cancel"
    on:click:button--secondary={()=>(openModal=false)}
>
    <TextInput labelText="Quota name"/>
    <NumberInput label="Default Quota value" value={0}/>
    
</Modal>