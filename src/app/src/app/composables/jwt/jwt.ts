import axios from "axios";

export const createNewAccessToken = async () => {
    let result: boolean
    const refreshToken: string | null = localStorage.getItem('refreshToken')
    if (refreshToken === null) {
        return
    }
    const params: URLSearchParams = new URLSearchParams();
    params.append('grant_type', 'refresh_token');
    params.append('refresh_token', refreshToken);

    try {
        // Set config defaults when creating the instance
        const instance = axios.create({
            baseURL: 'http://localhost:8000/',
            headers: {
                "Content-Type": "application/x-www-form-urlencoded",
            },
            params: params,
        });
        const response = await instance.post('/refresh-token')
        const accessToken: string = response.data.access_token
        const refreshToken: string = response.data.refresh_token

        localStorage.setItem('accessToken', accessToken)
        localStorage.setItem('refreshToken', refreshToken)
        result = true
    } catch (error) {
        console.log(error);
        result = false
    }

    return {result}
}