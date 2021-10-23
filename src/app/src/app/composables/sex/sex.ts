import axios from "axios";
import {onBeforeMount, readonly, ref} from "vue";
import {checkAuth} from "../auth/auth";
import {decodeJwt} from "../jwt/jwt";

type GetSexApiResponse = {
    config: object
    data: {
        sex_id: Number
        sex_code: Number
        sex_name: string
    }
    headers: object
    status: Number
    statusText: string
}

type Sex = {
    code: Number
    name: string
}

type sexLists = {
    sex_code: Number
    sex_id: Number
    sex_name: string
}

export const useSex = () => {
    const sexes = ref<Sex[] | null>(null)
    let useSexResult = ref<boolean>(false)
    onBeforeMount(async () => {
        checkAuth()
        const accessJwtToken: string | null = localStorage.getItem('accessToken')
        if (accessJwtToken === null) {
            useSexResult.value = false
            return
        }
        const accessToken: string = decodeJwt(accessJwtToken)
        try {
            // Set config defaults when creating the instance
            const instance = axios.create({
                baseURL: 'http://localhost:8000',
                headers: {
                    "Content-Type": "application/json",
                    'Authorization': 'Bearer ' + accessJwtToken,
                },
            });
            const response:GetSexApiResponse = await instance.get('/sex')
            const sexLists: Sex[] = response.data.data.sexes
                .map((sex: sexLists) => ({
                    code: sex.sex_code,
                    name: sex.sex_name,
                }))
            sexes.value = sexLists
            useSexResult.value = true
        } catch (error) {
            useSexResult.value = false
            console.log(error);
            return
        }
    })

    return {useSexResult: readonly(useSexResult), sexes: readonly(sexes)}
}