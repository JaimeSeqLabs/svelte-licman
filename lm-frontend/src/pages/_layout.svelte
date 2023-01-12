<script lang="ts">
    import { Header, HeaderNav, HeaderNavItem, Content } from "carbon-components-svelte";
    import { isActive } from "@roxi/routify";

    let navRoutes = [
        { text: "licenses", href: "/licenses"},
        { text: "products", href: "/products"},
        { text: "organizations", href: "/organizations"},
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
</Header>

<Content>
    <slot/>
</Content>