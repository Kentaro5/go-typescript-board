package Middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"api/utils"
)

// UserIDKey is used as a key for storing the UserID in context at middleware
type UserIDKey struct{}

// Define our struct
type AuthMiddleware struct {
	accessToken string
}

// Initialize it somewhere
func (amw *AuthMiddleware) Initialize() {
	amw.accessToken = ""
}

type GenericResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Middleware function, which will be called for each request
func (amw *AuthMiddleware) ValidateAccessToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if amw.accessToken != "" {
			//Initializeされた後に使われるので、ここでエラーを出す
		}
		reqToken := r.Header.Get("Authorization")
		splitToken := strings.Split(reqToken, "Bearer ")
		reqToken = splitToken[1]

		amw.accessToken = reqToken
		claims, err := utils.ValidateAccessToken(amw.accessToken)

		fmt.Println("jwtToken", claims.ExpiresAt)
		fmt.Println("Unix", time.Now().Unix())
		// アクセストークンを新しくするかどうかのフラグをセット
		//var refreshTokenContext context.Context = context.WithValue(r.Context(), "refreshAccessToken", false)
		var refreshTokenContext context.Context = context.WithValue(r.Context(), "refreshAccessToken", false)
		r = r.WithContext(refreshTokenContext)
		// もし、トークンの有効期限が切れていたら、新しくアクセストークンを作成するようにフラグをセット
		if claims.ExpiresAt < time.Now().Unix() {
			refreshTokenContext = context.WithValue(r.Context(), "refreshAccessToken", true)
			r = r.WithContext(refreshTokenContext)
		}

		if err != nil {
			// data.ToJSON(&GenericError{Error: err.Error()}, w)
			utils.ToJSON(&GenericResponse{Status: false, Message: "Authentication failed. Invalid token"}, w)
			return
		}

		// r.Context().Value()でユーザーIDを取得できるようにContextにセットしておく。
		// 参考になりそうな記事：https://medium.com/@agatan/http%E3%82%B5%E3%83%BC%E3%83%90%E3%81%A8context-context-7211433d11e6
		ctx := context.WithValue(r.Context(), UserIDKey{}, claims.UserID)
		r = r.WithContext(ctx)

		// 問題なければ、次に進める
		next.ServeHTTP(w, r)
	})
}
