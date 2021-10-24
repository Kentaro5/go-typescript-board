import axios from "axios";
import {readonly} from "vue";
import {checkAuth} from "../auth/auth";
import {decodeJwt} from "../jwt/jwt";

type GetCityApiResponse = {
    config: object
    data: {
        data: {
            cities: City[]
        }
    }
    headers: object
    status: Number
    statusText: string
}

type City = {
    id: Number
    code: Number
    prefCode: Number
    name: string
}

export const useCities = async (prefCode: Number | null) => {
    let changeCityResult:boolean = false
    let cities: City[];
    checkAuth()
    const accessJwtToken: string | null = localStorage.getItem('accessToken')
    if (accessJwtToken === null) {
        changeCityResult = false
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
        const response: GetPrefectureApiResponse = await instance.get('/city/' + prefCode)
        cities = response.data.data.cities
            .map((city: City) => ({
                id: city.id,
                code: city.code,
                prefCode: city.pref_code,
                name: city.name,
            }))
        changeCityResult = true
    } catch (error) {
        changeCityResult = false
        console.log(error);
        return
    }

    return {cities: readonly(cities), changeCityResult: readonly(changeCityResult)}
}