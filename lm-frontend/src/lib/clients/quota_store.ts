
import { writable, type Writable } from "svelte/store";

export type Quota = {name:string, value:string}

export const quotas:Writable<Quota[]> = writable([])

