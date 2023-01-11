<script lang="ts">
  import { Router } from "@roxi/routify";
  import { routes } from "../.routify/routes";

  import "carbon-components-svelte/css/all.css"
  import { Theme } from "carbon-components-svelte";
  
  
  export let theme: any = "g80" // "white" | "g10" | "g80" | "g90" | "g100"

</script>

<Theme bind:theme/>
<Router {routes}/>