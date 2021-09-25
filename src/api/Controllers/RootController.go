package Controllers

import (
	"context"
	"fmt"
	"net/http"
)

/**
 * 200レスポンスを返すだけのメソッド
 */
func Root(w http.ResponseWriter, request *http.Request) {
	header := w.Header()
	header.Set("Access-Control-Allow-Origin", "http://localhost:3000")
	header.Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	header.Set("Access-Control-Allow-Headers", "Authorization")

	// In case you don't have separate CORS middleware
	if request.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
}
