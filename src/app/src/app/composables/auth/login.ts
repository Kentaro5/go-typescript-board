import {reactive, ref, toRefs} from "vue";
import axios from "axios";

type loginInput = {
    loginEmail: string
    loginPassword: string
}

type LoginStatus =  'failed' | ''

export const useLogin = () => {
    const loginInput = reactive<loginInput>({
        loginEmail: '',
        loginPassword: '',
    })

    const loginStatus = ref<LoginStatus>('')

    const login = () => {
        const data = {
            email: loginInput.loginEmail,
            password: loginInput.loginPassword,
        }
        console.log(data);
        axios.post('http://localhost:8000/login', data).then(function (response) {
            const result = response.data
            console.log('FFFFFFFF');
            console.log(result);
            if (result.status === 200) {
                localStorage.setItem('accessToken', result.data.access_token)
                localStorage.setItem('refreshToken', result.data.refresh_token)
                localStorage.setItem('user', result.data.user)
                location.href = '/'
            } else if(result.status === 400) {
                loginStatus.value = 'failed'
            }
        })
    }

    const changeLoginStatus = (status: LoginStatus) => {
        loginStatus.value = status
    }

    return {...toRefs(loginInput), login, loginStatus, changeLoginStatus}
}