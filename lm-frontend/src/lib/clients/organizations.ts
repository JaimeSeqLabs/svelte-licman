import axios, { type AxiosResponse } from "axios"
import { LM_PRIVATE_API } from "./common"


export type ListAllOrgsResponse = {
	organizations: ListAllOrgsItem[],
}

export type ListAllOrgsItem = {
	id: string,
	name: string,
	contact: string,
	mail: string,
	country: string
}

export type DomainOrganization = {
	id: string,
	name: string,
	contact: string,
	mail: string,
	country: string
	licenses: string[],
	date_created: string,
	last_updated: string
}

export type DescribeOrgResponse = DomainOrganization

export type UpdateOrgRequest = DomainOrganization

export type CreateOrgRequest = DomainOrganization

const baseOrgClient = axios.create({
    baseURL: LM_PRIVATE_API + "/organizations",
    timeout: 1_000,
    headers: {'X-lm-client': 'org'}
})


export const listAllOrgs: () => Promise<AxiosResponse<ListAllOrgsResponse>> =
    () => baseOrgClient.get<ListAllOrgsResponse>("/")

export const describeOrg: (id: string) => Promise<AxiosResponse<DescribeOrgResponse>> =
    (id) => baseOrgClient.get<DescribeOrgResponse>(`/${id}`)

export const updateOrg: (id: string) => Promise<AxiosResponse<void>> =
    (id) => baseOrgClient.put(`/${id}`)

export const createOrg: (req: CreateOrgRequest) => Promise<AxiosResponse<void>> = 
    (req) => baseOrgClient.post("/", req)

export const deleteOrg: (id: string) => Promise<AxiosResponse<void>> = 
    (id) => baseOrgClient.patch(`/delete/${id}`)

export const restoreOrg: (id: string) => Promise<AxiosResponse<void>> = 
    (id) => baseOrgClient.patch(`/restore/${id}`)