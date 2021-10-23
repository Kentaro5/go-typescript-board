import axios from "axios";
import {onBeforeMount, readonly, ref} from "vue";
import {checkAuth} from "../auth/auth";
import {decodeJwt} from "../jwt/jwt";

type GetUserApiResponse = {
    config: object
    data: {
        user_name: string
        email: string
        sex: string
        sex_code: Number
        prefecture: string
        prefecture_code: Number
        city: string
        city_code: Number
        ward: string | null
        war_code: Number | null
        created_at: string
    }
    headers: object
    status: Number
    statusText: string
}

type User = {
    name: string
    email: string
    sex: string
    sexCode: Number
    prefecture: string
    prefectureCode: Number
    city: string
    cityCode: Number
    ward: string | null
    wardCode: Number | null
    registeredDate: string
}

export const useUser = () => {
    const user = ref<User | null>(null)
    let useUserResult = ref<boolean>(false)
    onBeforeMount(async () => {
        checkAuth()
        const accessJwtToken: string | null = localStorage.getItem('accessToken')
        if (accessJwtToken === null) {
            useUserResult.value = false
            return
        }
        const accessToken:string = decodeJwt(accessJwtToken)
        const userId:Number = accessToken.UserID
        try {
            // Set config defaults when creating the instance
            const instance = axios.create({
                baseURL: 'http://localhost:8000',
                headers: {
                    "Content-Type": "application/json",
                    'Authorization': 'Bearer ' + accessJwtToken,
                },
            });

            const response:GetUserApiResponse = await instance.get('/user/' + userId)
            user.value = {
                name: response.data.data.user_name,
                email: response.data.data.email,
                sex: response.data.data.sex,
                sexCode: response.data.data.sex_code,
                prefecture: response.data.data.prefecture,
                prefectureCode: response.data.data.prefecture_code,
                city: response.data.data.city,
                cityCode: response.data.data.city_code,
                ward: response.data.data.ward,
                wardCode: response.data.data.ward_code,
                registeredDate: response.data.data.created_at,
            }
            useUserResult.value = true
        } catch (error) {
            useUserResult.value = false
            return
        }
    })

    return {useUserResult: readonly(useUserResult), user: readonly(user)}
}