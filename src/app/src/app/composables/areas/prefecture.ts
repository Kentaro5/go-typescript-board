import axios from "axios";
import {onBeforeMount, readonly, ref, watch} from "vue";
import {checkAuth} from "../auth/auth";
import {decodeJwt} from "../jwt/jwt";

type GetPrefectureApiResponse = {
    config: object
    data: {
        data: {
            prefectures: Prefecture[]
        }
    }
    headers: object
    status: Number
    statusText: string
}

type Prefecture = {
    id: Number
    code: Number
    name: string
}

export const usePrefectures = () => {
    const prefectures = ref<Prefecture[] | null>(null)
    let usePrefecturesResult = ref<boolean>(false)
    let prefectureIndex:Number = 0
    onBeforeMount(async () => {
        checkAuth()
        const accessJwtToken: string | null = localStorage.getItem('accessToken')
        if (accessJwtToken === null) {
            usePrefecturesResult.value = false
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
            const response:GetPrefectureApiResponse = await instance.get('/prefecture')
            const prefectureLists: Prefecture[] = response.data.data.prefectures
                .map((prefecture: Prefecture) => ({
                    id: prefecture.id,
                    code: prefecture.code,
                    name: prefecture.name,
                }))
            prefectures.value = prefectureLists
            usePrefecturesResult.value = true
        } catch (error) {
            usePrefecturesResult.value = false
            console.log(error);
            return
        }
    })

    return {usePrefecturesResult: readonly(usePrefecturesResult), prefectures: readonly(prefectures), prefectureIndex}
}