import axios, { type AxiosResponse } from "axios"
import { LM_PRIVATE_API } from "./common"

export type CreateProductRequest = {
	sku: string,
	name: string,
	install_instructions: string
}

export type ListAllProductsResponse = {
	products: ListAllProductsItem[]
}

export type DescribeProductResponse = {
	id: string,
	sku: string,
	name: string,
	install_instructions: string,
	license_count: number,
	
	date_created: string,
	last_updated: string,
}

export type ListAllProductsItem = DescribeProductResponse

export type UpdateProductRequest = CreateProductRequest

const baseProdClient = axios.create({
    baseURL: LM_PRIVATE_API + "/products",
    timeout: 1_000,
    headers: {'X-lm-client': 'prod'}
})

export const listAllProducts: () => Promise<AxiosResponse<ListAllProductsResponse>> =
    () => baseProdClient.get<ListAllProductsResponse>("/")

export const createNewProduct: (req: CreateProductRequest) => Promise<AxiosResponse<void>> =
    (req) => baseProdClient.post("/", req)

export const describeProduct: (id: string) => Promise<AxiosResponse<DescribeProductResponse>> =
    (id) => baseProdClient.get<DescribeProductResponse>(`/${id}`)

export const updateProduct: (req: {id: string, update: UpdateProductRequest}) => Promise<AxiosResponse<void>> =
    ({id, update}) => baseProdClient.put(`/${id}`, update)

export const deleteProduct: (id: string) => Promise<AxiosResponse<void>> =
    (id) => baseProdClient.patch(`/delete/${id}`)

export const restoreProduct: (id: string) => Promise<AxiosResponse<void>> =
    (id) => baseProdClient.patch(`/restore/${id}`)
