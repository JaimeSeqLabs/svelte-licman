<script lang="ts">
    import { Button, TextArea } from "carbon-components-svelte";
    import { Unlocked } from "carbon-icons-svelte";

</script>

<TextArea
    labelText="Decode License"
    placeholder="Provide encoded license certificate string"
/>
<Button 
    iconDescription={"Decode"}
    icon={Unlocked}
>
    Decode
</Button>