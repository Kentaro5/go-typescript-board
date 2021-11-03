import axios from "axios";

export const createNewAccessToken = async () => {
    let result: boolean
    const refreshToken: string | null = localStorage.getItem('refreshToken')
    if (refreshToken === null) {
        return false
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
        const accessToken: string = response.data.data.access_token
        const refreshToken: string = response.data.data.refresh_token
        localStorage.setItem('accessToken', accessToken)
        localStorage.setItem('refreshToken', refreshToken)

        result = true
    } catch (error) {
        console.log(error);
        result = false
    }

    return result
}

/**
 * Jwtのコードをデコードする。
 * @param token
 */
export const decodeJwt = (token: string) => {
    const base64Url = token.split('.')[1];
    const base64 = base64Url.replace(/-/g, '+').replace(/_/g, '/');
    const jsonPayload = decodeURIComponent(atob(base64).split('').map(function(c) {
        return '%' + ('00' + c.charCodeAt(0).toString(16)).slice(-2);
    }).join(''));

    return JSON.parse(jsonPayload);
}

/**
 * アクセストークンの有効期限をチェックする
 * 有効期限が切れていた場合は、trueを返す
 * @param expiredTime
 */
export const isExpiredAccessToken = (expiredTime: number) => {
    // Dateオブジェクトを作成
    const dateObject: Date = new Date() ;
    // 一度ミリ秒単位で、UNIXタイムスタンプを取得する
    const milliSecondTimeStamp: number = dateObject.getTime() ;
    // そのあと、秒単位UNIXタイムスタンプを生成
    const secondTimeStamp: number = Math.floor( milliSecondTimeStamp / 1000 ) ;

    if (expiredTime > secondTimeStamp) {
        return false
    }

    return true
}
