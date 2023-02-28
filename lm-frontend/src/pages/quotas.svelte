<script lang="ts">
    import { Modal, NumberInput, TextInput, TextInputSkeleton } from "carbon-components-svelte";
    import { createQuota, listAllQuotas } from "../lib/clients/quotas";
    import QuotasList from "../lib/components/quotas/QuotasList.svelte";

    let openModal = false

    let qName = ""
    let qValue: number = 0

    function handleCreateQuota() {
        
        if (qName != "") {
            console.log("creating quota");
            createQuota(qName, qValue.toString())
            console.log(listAllQuotas());
        }
        qName = ""
        qValue = 0        
        openModal = false   
    }

</script>

<QuotasList onCreateButton={()=>{ openModal = true }}/>

<Modal
    bind:open={openModal}
    modalHeading="Create Quota"
    primaryButtonText="Save"
    secondaryButtonText="Cancel"
    on:click:button--secondary={()=>(openModal=false)}
    on:click:button--primary={handleCreateQuota}
>
    <TextInput labelText="Quota name" bind:value={qName}/>
    <NumberInput label="Default Quota value" bind:value={qValue}/>
    
</Modal>