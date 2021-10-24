import axios from "axios";
import {readonly} from "vue";
import {checkAuth} from "../auth/auth";
import {decodeJwt} from "../jwt/jwt";

type GetWardiResponse = {
    config: object
    data: {
        data: {
            wards: Ward[]
        }
    }
    headers: object
    status: Number
    statusText: string
}

type Ward = {
    id: Number
    code: Number
    cityCode: Number
    name: string
}

export const useWards = async (cityCode: Number | null) => {
    let changeWardResult:boolean = false
    let wards: Ward[] | [];
    checkAuth()
    const accessJwtToken: string | null = localStorage.getItem('accessToken')
    if (accessJwtToken === null) {
        changeWardResult = false
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
        const response: GetWardiResponse = await instance.get('/ward/' + cityCode)
        wards = []
        if (response.data.data.wards !== null) {
            wards = response.data.data.wards
                .map((ward: Ward) => ({
                    id: ward.id,
                    code: ward.code,
                    cityCode: ward.city_code,
                    name: ward.name,
                }))
            changeWardResult = true
        }
    } catch (error) {
        changeWardResult = false
        console.log(error);
        return
    }

    return {wards: wards, changeWardResult: readonly(changeWardResult)}
}