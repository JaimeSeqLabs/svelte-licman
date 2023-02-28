
// TODO: this is a mock

import { quotas, type Quota } from "./quota_store"

let currentQuotas:Quota[] = []

quotas.subscribe(q => currentQuotas = q)

quotas.set([
    { name: "maxUsersPerWsp", value: "10"},
    { name: "maxWspPerOrg", value: "15"}
])

export const listAllQuotas = () => currentQuotas

export const createQuota = (name:string, value: string) => {
    quotas.update(q => [...q, {name, value}])
}

export const getQuota = (name:string): string | undefined => {

    return currentQuotas?.find(q => q.name == name).value

}

export const updateQuota = (name: string, value: string) => {

    quotas.update(qs => {
        for (const quota of qs) {
            if (quota.name == name) {
                quota.value = value
            }
        }
        return qs
    })

}

export const deleteQuota = (name: string) => {
    quotas.update(qs => qs.filter(q => q.name != name))
}