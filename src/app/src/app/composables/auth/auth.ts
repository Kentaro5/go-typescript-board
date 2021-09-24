import axios from "axios";

export const checkAuth = async () => {
    let result: boolean = false
    try {
        const options = {
            withCredentials: true,
        };
        const response = await axios.get('http://localhost:8000/', options)
        result = true
        console.log(response);
    } catch (error) {
        result = false
        console.log(error);
    }

    if (!result) {
        location.href = 'login'
        return
    }

    return {result}
}
