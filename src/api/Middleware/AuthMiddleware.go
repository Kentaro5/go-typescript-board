package Middleware

import (
	"context"
	"fmt"
	"net/http"

	"api/utils"
)

// UserIDKey is used as a key for storing the UserID in context at middleware
type UserIDKey struct{}

// Define our struct
type AuthMiddleware struct {
	cookieAccessToken string
}

// Initialize it somewhere
func (amw *AuthMiddleware) Initialize() {
	amw.cookieAccessToken = ""
}

type GenericResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Middleware function, which will be called for each request
func (amw *AuthMiddleware) ValidateAccessToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if amw.cookieAccessToken != "" {
			//Initializeされた後に使われるので、ここでエラーを出す
		}

		// クッキーからアクセストークンを取得
		cookie, err := utils.GetToken(r)
		if err != nil {
			http.Error(w, "Cookie not found", http.StatusForbidden)
			utils.ToJSON(&GenericResponse{Status: false, Message: "Cookie not found"}, w)
			return
		}
		amw.cookieAccessToken = cookie
		fmt.Println(amw.cookieAccessToken)
		userID, err := utils.ValidateAccessToken(amw.cookieAccessToken)
		if err != nil {
			// data.ToJSON(&GenericError{Error: err.Error()}, w)
			utils.ToJSON(&GenericResponse{Status: false, Message: "Authentication failed. Invalid token"}, w)
			return
		}

		// r.Context().Value()でユーザーIDを取得できるようにContextにセットしておく。
		// 参考になりそうな記事：https://medium.com/@agatan/http%E3%82%B5%E3%83%BC%E3%83%90%E3%81%A8context-context-7211433d11e6
		ctx := context.WithValue(r.Context(), UserIDKey{}, userID)
		r = r.WithContext(ctx)

		// 問題なければ、次に進める
		next.ServeHTTP(w, r)
	})
}
