<script lang="ts">
    
    import { Button, FluidForm, PasswordInput, TextInput, Tile } from "carbon-components-svelte";
    import { Login } from "carbon-icons-svelte";
    import { goto } from "@roxi/routify";
    import { doLoginWith, isLoggedIn } from "../../lib/clients/login";
    import logo from "/seqera_logo_white.png"
    import { onMount } from "svelte";

    let mail: string
    let pwd: string
    
    let errorMsg: string = ""
    const showErr = (msg: string) => {
        errorMsg = msg
        setTimeout(() => errorMsg = "", 5_000);
    }

    const handleClick = () => {
        doLoginWith(mail, pwd)
        .then((res) => {

            if (res.status != 200) {
                showErr("Wrong credentials")
            } else {
                $goto("/")
            }

        })
        .catch(_ => {
            showErr("Wrong credentials")
        })
    }

    onMount(async () => {
        console.log("check login from login page");
        
        if (await isLoggedIn()) {
            $goto("/licenses")
        }
    })

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
                labelText="User mail"
                placeholder="Enter user mail..."
                bind:value={mail}
            />
            <PasswordInput
                required
                type="password"
                labelText="Password"
                placeholder="Enter password..."
                bind:value={pwd}
            />
        </FluidForm>
        
        <Tile>
            <div class="flex justify-end">

                <Button 
                    size="field" 
                    icon={Login} 
                    iconDescription="Login"
                    on:click={handleClick}
                >
                    Login
                </Button>
            
            </div>
        </Tile>

        {#if errorMsg != ""}
            <Tile>
                <p class="
                    bg-gray-500 
                    font-bold text-sm
                    border-red-500 border-2 border-solid
                    rounded p-2
                ">
                    {errorMsg}
                </p>
            </Tile>
        {/if}
        

    </div>
</div>