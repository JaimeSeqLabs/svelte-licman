import type { AxiosResponse } from "axios";

export const LM_PRIVATE_API = "/lm/api"
export const LM_PUBLIC_API = "/lm"

export const errWrap = <ResT>(f: (req?: any)=>Promise<AxiosResponse<ResT, any>>) => 
    (req?: any) => f(req)
        .then(res => res.data)
        .catch(err => {
            console.error(err.message, err.status);
            return {} as ResT
        })