import {createNewAccessToken} from "../jwt/jwt";
import {onBeforeMount} from 'vue'

export const checkAuth = () => {
    onBeforeMount(async () => {
        const accessJwtToken: string | null = localStorage.getItem('accessToken')
        const refreshJwtToken: string | null = localStorage.getItem('refreshToken')
        if (!accessJwtToken || accessJwtToken === "undefined" || !refreshJwtToken || refreshJwtToken === "undefined") {
            // 各トークンが存在しない場合は、ログイン画面へリダイレクト
            location.href = '/login'
            return
        }
        const accessToken = decodeJwt(accessJwtToken)
        if (isExpiredAccessToken(accessToken.exp)) {
            // トークンの有効期限が切れていた場合は、新しくアクセストークンを取り直す
            const result = await createNewAccessToken()
            if (!result) {
                // もしトークンが取れなかった場合は、各ストレージの値を空にした上で、ログイン画面へリダイレクト
                localStorage.removeItem('accessToken')
                localStorage.removeItem('refreshToken')
                location.href = '/login'
                return
            }
        }
    })
}

/**
 * Jwtのコードをデコードする。
 * @param token
 */
const decodeJwt = (token: string) => {
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
const isExpiredAccessToken = (expiredTime: number) => {
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
