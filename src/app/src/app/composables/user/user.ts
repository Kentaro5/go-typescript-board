import axios from "axios";
import {onBeforeMount, readonly, ref, reactive} from "vue";
import {checkAuth} from "../auth/auth";
import {decodeJwt} from "../jwt/jwt";
import {useCities} from "../areas/city";
import {useWards} from "../areas/ward";

type GetUserApiResponse = {
    config: object
    data: {
        user_name: string
        email: string
        sex: string
        sex_code: Number
        prefecture: string
        pref_code: Number
        city: string
        city_code: Number
        ward: string | null
        ward_code: Number | null
        created_at: string
    }
    headers: object
    status: Number
    statusText: string
}

type PatchUserApiResponse = {
    config: object
    data: ""
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

type Updated = {
    status: 'updated' | 'error' | 'unupdated'
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
                prefectureCode: response.data.data.pref_code,
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

export const useEditUser = () => {
    const user = ref<User | null>(null)
    const userSexCode = ref<Number | null>(null)
    const userPrefectureCode = ref<Number | string>('')
    const userCityCode = ref<Number | string>('')
    const userWardCode = ref<Number | string>('')
    const updated = reactive<Updated>({
        status: 'unupdated'
    })

    const changeUpdatedStatus = (status) => {
        updated.status = status
    }

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
                prefectureCode: response.data.data.pref_code,
                prefectureLists: [],
                city: response.data.data.city,
                cityLists: [],
                cityCode: response.data.data.city_code,
                ward: response.data.data.ward,
                wardLists: [],
                wardCode: response.data.data.ward_code,
                registeredDate: response.data.data.created_at,
            }
            useUserResult.value = true
            userSexCode.value = user.value.sexCode
            userPrefectureCode.value = user.value.prefectureCode
            userCityCode.value = user.value.cityCode
            userWardCode.value = user.value.wardCode

            console.log(user.value.prefectureCode);
            const {cities} = await useCities(user.value.prefectureCode)
            user.value.cityLists = cities

            if (user.value.wardCode !== null) {
                const {wards} = await useWards(user.value.cityCode)
                user.value.wardLists = wards
                console.log(wards);
            }
        } catch (error) {
            useUserResult.value = false
            return
        }
    })

    const updateUserInfo = async () => {
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

            const data = {
                'name': user.value.name,
                'sex_code': userSexCode.value,
                'email': user.value.email,
                'pref_code': userPrefectureCode.value,
                'city_code': userCityCode.value,
                'ward_code': userWardCode.value,
            }

            const response: PatchUserApiResponse = await instance.patch('/user/' + userId, data)

            if (response.status === 200) {
                updated.status = 'updated'
            } else {
                updated.status = 'error'
            }
        } catch (error) {
            updated.status = 'error'
            return
        }
    }

    return {
        useUserResult: readonly(useUserResult),
        user,
        updateUserInfo,
        changeUpdatedStatus,
        userSexCode,
        userPrefectureCode,
        userCityCode,
        userWardCode,
        updated,
    }
}
