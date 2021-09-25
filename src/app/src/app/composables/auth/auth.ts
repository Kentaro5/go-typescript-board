import axios from "axios";

export const checkAuth = async () => {
    let result: boolean = false
    try {
        // Set config defaults when creating the instance
        const instance = axios.create({
            baseURL: 'http://localhost:8000/',
            headers: {
                "Authorization": `Bearer ${localStorage.getItem('accessToken')}`,
                "Content-Type": "application/json",
            },
            data: {}
        });
        const response = await instance.get('/')
        result = true
        console.log(response);
    } catch (error) {
        result = false
        console.log(error);
    }

    if (!result) {
        // location.href = 'login'
        // return
    }

    return {result}
}
