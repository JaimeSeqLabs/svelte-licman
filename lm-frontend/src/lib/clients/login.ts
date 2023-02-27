import axios, { type AxiosResponse } from "axios";
import { errWrap, LM_PRIVATE_API, LM_PUBLIC_API } from "./common";
import { goto } from "@roxi/routify";


export type LoginCredentials = {
    user?: string,
    mail?: string,
    password_hash: string,
}

export type JWTResponse = {
    access_token: string,
}

const sha256 = async (msg: string): Promise<string> => {

    // encode as UTF-8
    const msgBuffer = new TextEncoder().encode(msg);                    

    // hash the message
    const hashBuffer = await crypto.subtle.digest('SHA-256', msgBuffer);

    // convert ArrayBuffer to Array
    const hashArray = Array.from(new Uint8Array(hashBuffer));

    // convert bytes to hex string                  
    const hashHex = hashArray.map(b => b.toString(16).padStart(2, '0')).join('');
    return hashHex;
}

const baseLoginClient = axios.create({
    baseURL: LM_PUBLIC_API,
    timeout: 1_000,
    headers: {'X-lm-client': 'login'}
})

export const doLogin: (req: LoginCredentials) => Promise<AxiosResponse<JWTResponse>> = 
        (req) => baseLoginClient.post<JWTResponse>("/login", req)

export const doLoginWith: (mail: string, password: string) => Promise<AxiosResponse<JWTResponse>> = 
    (mail: string, password: string) => 
        sha256(password)
        .then(password_hash => doLogin({ mail, password_hash}))

export const isLoggedIn: () => Promise<boolean> = 
    () => axios.get(LM_PRIVATE_API + "/is_admin")
        .then(res => res.status == 200)
        .catch(err => false)

export const logOut = () => {
    document.cookie = "jwt=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;";
    $goto("/login")
}