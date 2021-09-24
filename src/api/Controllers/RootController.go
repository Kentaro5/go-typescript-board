package Controllers

import (
	"net/http"
)

/**
 * 200レスポンスを返すだけのメソッド
 */
func Root(w http.ResponseWriter, request *http.Request) {
	header := w.Header()
	header.Set("Content-Type", "application/json")
	header.Set("Access-Control-Allow-Credentials", "true")
	header.Set("Access-Control-Allow-Headers", "Content-Type, withCredentials")
	header.Set("Access-Control-Allow-Origin", "http://localhost:3000")
	header.Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	// In case you don't have separate CORS middleware
	if request.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
}
