import axios, { type AxiosResponse } from "axios";
import { LM_PRIVATE_API, errWrap } from "./common";


export type LicenseStatus = "archived" | "suspended" | "expired" | "active";

export type CreateLicenseRequest = {
	
	features: string,
	status: LicenseStatus,
	version: string,
	
	note: string,
	contact: string,
	mail: string,
	
	// NOTE: skus not IDs
	product_skus: string,
	// NOTE: name not ID
	organization_name: string,

	quotas: Map<string, string>,

	secret: string,

	expiration_date: string,
	activation_date: string
}

export type UpdateLicenseRequest = {

	license: DomainLicense,
	product_ids: string[],// product IDs
	quotas: Map<string, string>,
	
}

export type ListAllLicensesResponse = {
	licenses: ListAllLicensesItem[],
}

export type ListAllLicensesItem = DomainLicense

export type DomainLicense = {
    id: string,

	features: string,
	status: LicenseStatus, // archived, suspended, expired, active
	version: string,
	
	note: string,
	contact: string,
	mail: string,
	
    product_ids: string[],
	organization_id: string,
	quotas: Map<string, string>,

	expiration_date: string,
	activation_date: string,
	date_created: string,
	last_updated: string,
	last_accessed: string,

	access_count: number,
}

export type DescribeLicenseResponse = {
	license: DomainLicense,
	quotas: Map<string, string>,
}

export type DecodeLicenseRequest = {
	encoded: string,
}

export type DescribeLicenseStatusResponse = {
	
	id: string,
	organization_id: string,
	mail: string,
	activation_date: string,
	expiration_date: string,
	status: LicenseStatus, // archived, suspended, expired, active
	contact: string,
	product_ids: string[],
	quotas: Map<string, string>,

}

const baseLicenseClient = axios.create({
    baseURL: LM_PRIVATE_API + "/licenses",
    timeout: 1_000,
    headers: {'X-lm-client': 'license'}
})

export const listAllLicenses: () => Promise<ListAllLicensesResponse> = 
    errWrap(
        () => baseLicenseClient.get<ListAllLicensesResponse>("/")
    )

export const createNewLicense: (req: CreateLicenseRequest) => Promise<DomainLicense> = 
    errWrap(
        (req) => baseLicenseClient.post<DomainLicense>("/", req)
    )

export const decodeLicense: (req: DecodeLicenseRequest) => Promise<DescribeLicenseResponse> = 
    errWrap(
        (req) => baseLicenseClient.post<DescribeLicenseResponse>("/decode/", req)
    )

export const deleteLicense: (id: string) => Promise<void> = 
    errWrap(
        (id) => baseLicenseClient.patch(`/delete/${id}`)
    )

export const restoreLicense: (id: string) => Promise<void> = 
    errWrap(
        (id) => baseLicenseClient.patch(`/restore/${id}`)
    )

export const downloadLicense: (id: string) => Promise<void> = 
    errWrap(
        (id) => baseLicenseClient.get(`/download/${id}`)
    )

export const expireLicense: (id: string) => Promise<void> = 
    errWrap(
        (id) => baseLicenseClient.patch(`/expire/${id}`)
    )

export const getStatusForLicense: (id: string) => Promise<DescribeLicenseStatusResponse> = 
    errWrap(
        (id) => baseLicenseClient.get(`/status/${id}`)
    )

export const describeLicense: (id: string) => Promise<DescribeLicenseResponse> = 
    errWrap(
        (id: string) => baseLicenseClient.get<DescribeLicenseResponse>(`/${id}`)
    )

export const updateLicense: (req: {id: string, update: UpdateLicenseRequest}) => Promise<DescribeLicenseResponse> = 
    errWrap(
        ({id, update}) => baseLicenseClient.put<DescribeLicenseResponse>(`/${id}`, update)
    )

