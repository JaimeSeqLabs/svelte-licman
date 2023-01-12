<script lang="ts">
    import { UserAvatarFilledAlt } from "carbon-icons-svelte";
    import { isActive } from "@roxi/routify";
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

    let navRoutes = [
        { text: "licenses", href: "/licenses"},
        { text: "organizations", href: "/organizations"},
        { text: "products", href: "/products"},
        { text: "decoder", href: "/decoder"},
        { text: "quotas", href: "/quotas"}
    ]

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
                <HeaderPanelLink>Logout</HeaderPanelLink>
            </HeaderPanelLinks>
        </HeaderAction>
    </HeaderUtilities>

</Header>

<Content>
    <slot/>
</Content>