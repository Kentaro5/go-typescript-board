import {onBeforeMount, reactive, toRefs} from "vue";
import {checkAuth} from "../auth/auth";
import {decodeJwt} from "../jwt/jwt";
import axios from "axios";

type ChangeUserPasswordApiResponse = {
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

type EditPasswords = {
    oldPassword: string
    oldConfirmPassword: string
    newPassword: string
    updated: {
        status: string
    }
}

export const useEditPassword = () => {
    const editPassword = reactive<EditPasswords>({
        oldPassword: '',
        oldConfirmPassword: '',
        newPassword: '',
        updated: {
            status: ''
        }
    })

    onBeforeMount(async () => {
        checkAuth()
    })

    const changeUpdatedStatus = (status) => {
        editPassword.updated.status = status
    }

    const updatePassword = async () => {
        checkAuth()
        const accessJwtToken: string | null = localStorage.getItem('accessToken')
        const refreshJwtToken: string | null = localStorage.getItem('refreshToken')
        if (accessJwtToken === null) {
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
                'old_password': editPassword.oldPassword,
                'new_password': editPassword.newPassword,
                'refresh_token': refreshJwtToken,
                'grant_type': 'refresh_token',
            }

            const response: ChangeUserPasswordApiResponse = await instance.patch('/user/' + userId + '/changePassword/', data)

            if (response.status === 200) {
                editPassword.updated.status = 'updated'
                const accessToken: string = response.data.data.access_token
                const refreshToken: string = response.data.data.refresh_token
                localStorage.removeItem('accessToken')
                localStorage.removeItem('refreshToken')
                localStorage.setItem('accessToken', accessToken)
                localStorage.setItem('refreshToken', refreshToken)
            } else {
                editPassword.updated.status = 'error'
            }
        } catch (error) {
            editPassword.updated.status = 'error'
            return
        }
    }

    return {...toRefs(editPassword), updatePassword, changeUpdatedStatus}
}