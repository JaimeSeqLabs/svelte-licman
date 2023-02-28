<script lang="ts">
    import { UserAvatarFilledAlt } from "carbon-icons-svelte";
    import { isActive, goto } from "@roxi/routify";
    import { 
        Header, 
        Content, 
        HeaderNav, 
        HeaderNavItem, 
        HeaderUtilities, 
        HeaderAction, 
        HeaderPanelLink, 
        HeaderPanelLinks 
    } from "carbon-components-svelte";
    import { isLoggedIn, dropJwt } from "../lib/clients/login";
    import { onMount } from "svelte";

    let navRoutes = [
        { text: "licenses", href: "/licenses"},
        { text: "organizations", href: "/organizations"},
        { text: "products", href: "/products"},
        { text: "decoder", href: "/decoder"},
        { text: "quotas", href: "/quotas"}
    ]

    function handleLogout() {
        dropJwt()
        $goto("/login")
    }

    // login guard
    onMount(async () => {
        console.log("check login from main layout");
        
        if (!await isLoggedIn()) {
            $goto("/login")
        }
    })

</script>

<Header 
    company="Seqera Labs"
    platformName="License Manager"
    href="/"
>
    <HeaderNav>
        {#each navRoutes as item}
            <HeaderNavItem {...item} isSelected={$isActive(item.href)}></HeaderNavItem>
        {/each}
    </HeaderNav>

    <HeaderUtilities>
        <HeaderAction
            icon={UserAvatarFilledAlt}
            closeIcon={UserAvatarFilledAlt}
        >
            <HeaderPanelLinks>
                <HeaderPanelLink on:click={handleLogout}>
                    Logout
                </HeaderPanelLink>
            </HeaderPanelLinks>
        </HeaderAction>
    </HeaderUtilities>

</Header>

<Content>
    <slot/>
</Content>