<script lang="ts">
    import { Button, FluidForm, PasswordInput, TextInput, Tile } from "carbon-components-svelte";
    import { Login } from "carbon-icons-svelte";
    import logo from "/seqera_logo_white.png"

</script>

<style>
    img {
        @apply p-4;
        opacity: 1;
        background-color: #0f62fe;
    }
</style>

<div class="h-full flex flex-wrap justify-center content-center">
    <div class="w-96">

        <img src={logo} alt="">

        <Tile class="text-xl font-bold">Staff access</Tile>
            
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
        
        <Tile>
            <div class="flex justify-end">
                <Button size="field" icon={Login} iconDescription="Login">Login</Button>
            </div>
        </Tile>

    </div>
</div>