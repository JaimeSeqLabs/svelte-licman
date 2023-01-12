<script lang="ts">
    import { Column, FluidForm, Grid, PasswordInput, Row, TextInput } from "carbon-components-svelte";
    import SeqeraLogo from "../../lib/components/SeqeraLogo.svelte";


</script>

<div class="h-full flex flex-wrap justify-center content-center">
    <div class="w-2/5">
        <SeqeraLogo/>
        <FluidForm>
            <TextInput
                required
                labelText="User name"
                placeholder="Enter user name..."
            />
            <PasswordInput
                required
                type="password"
                labelText="Password"
                placeholder="Enter password..."
            />
        </FluidForm>
    </div>
</div>