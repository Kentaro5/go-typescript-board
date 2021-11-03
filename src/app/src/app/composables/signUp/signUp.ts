import {reactive, ref, toRefs, onBeforeMount, readonly} from "vue";
import axios from "axios";

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

type GetWardResponse = {
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

type Prefecture = {
    id: Number
    code: Number
    name: string
}

type City = {
    id: Number
    code: Number
    prefCode: Number
    name: string
}

type Ward = {
    id: Number
    code: Number
    cityCode: Number
    name: string
}

type SignUpUser = {
    name: string
    sexCode: Number
    password: string
    email: string
    prefCode: Number
    cityCode: Number
    wardCode: Number | null
}

export const useSingUp = () => {
    const signUpUser = reactive<SignUpUser>({
        name: '',
        sexCode: 0,
        email: '',
        password: '',
        prefCode: 0,
        cityCode: 0,
        wardCode: null,
    })
    const signUpResult = ref<boolean>(false)

    const sexes = ref<Sex[] | null>(null)
    let useSexResult = ref<boolean>(false)

    const prefectures = ref<Prefecture[] | null>(null)
    let usePrefecturesResult = ref<boolean>(false)

    onBeforeMount(async () => {
        await getPrefectures()
        await getSexes()
    })

    const getPrefectures = async () => {
        try {
            // Set config defaults when creating the instance
            const instance = axios.create({
                baseURL: 'http://localhost:8000',
            });
            const response:GetPrefectureApiResponse = await instance.get('/get/prefecture')
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
    }

    const getSexes = async () => {
        try {
            // Set config defaults when creating the instance
            const instance = axios.create({
                baseURL: 'http://localhost:8000',
                headers: {
                    "Content-Type": "application/json",
                },
            });
            const response:GetSexApiResponse = await instance.get('/get/sex')
            console.log('GetSexApiResponse');
            console.log(response);
            console.log('GetSexApiResponse');
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
    }

    const signUp = () => {
        const data = {
            name: signUpUser.name,
            sex_code: signUpUser.sexCode,
            password: signUpUser.password,
            email: signUpUser.email,
            pref_code: signUpUser.prefCode,
            city_code: signUpUser.cityCode,
            ward_code: signUpUser.wardCode,
        }
        axios.post('http://localhost:8000/signUp', data).then(function (response) {
            const result = response.data
            if (result.status === 200) {
                signUpResult.value = true
            } else if(result.status === 400) {
                signUpResult.value = false
            }
        })
    }

    const changeUpdatedStatus = (status: boolean) => {
        useSexResult.value = status
    }

    return {
        changeUpdatedStatus,
        signUpResult,
        signUp,
        ...toRefs(signUpUser),
        usePrefecturesResult: readonly(usePrefecturesResult),
        prefectures: readonly(prefectures),
        useSexResult: readonly(useSexResult),
        sexes: readonly(sexes),
    }
}

export const getCities = async (prefCode: Number | null) => {
    let changeCityResult:boolean = false
    let cities: City[];

    try {
        // Set config defaults when creating the instance
        const instance = axios.create({
            baseURL: 'http://localhost:8000',
            headers: {
                "Content-Type": "application/json",
            },
        });
        const response: GetPrefectureApiResponse = await instance.get('/get/city/' + prefCode)
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

    return {
        cities: cities,
        changeCityResult: readonly(changeCityResult),
    }
}

export const getWards = async (cityCode: Number | null) => {
    let changeWardResult:boolean = false
    let wards: Ward[] | [];

    try {
        // Set config defaults when creating the instance
        const instance = axios.create({
            baseURL: 'http://localhost:8000',
            headers: {
                "Content-Type": "application/json",
            },
        });
        const response: GetWardResponse = await instance.get('/get/ward/' + cityCode)
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

    return {
        wards: wards,
        changeWardResult: readonly(changeWardResult),
    }
}
