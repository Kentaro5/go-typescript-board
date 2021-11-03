import {createNewAccessToken, decodeJwt, isExpiredAccessToken} from "../jwt/jwt";
import {onBeforeMount} from 'vue'

export const checkAuth = () => {
    // TODO: onBeforeMountで使うのが正しくないため、あとで修正
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
