package Controllers

import (
	"api/db"
	"api/infrastructure/disabledRefreshTokenRepository"
	"api/infrastructure/userRepositopry"
	"api/utils"
	"fmt"
	"log"
	"net/http"
)

type NewTokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

/**
 * 新しくリフレッシュトークンとアクセストークンを作成して返す
 */
func CreateAccessTokenByRefreshToken(w http.ResponseWriter, request *http.Request) {
	header := w.Header()
	header.Set("Content-Type", "application/x-www-form-urlencoded")
	header.Set("Access-Control-Allow-Origin", "http://localhost:3000")
	header.Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	header.Set("Access-Control-Allow-Headers", "Authorization")

	// In case you don't have separate CORS middleware
	if request.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if err := request.ParseForm(); err != nil {
		fmt.Println("errorだよ")
	}

	grantType := request.Form.Get("grant_type")
	refreshToken := request.Form.Get("refresh_token")
	fmt.Println("grantType", grantType)
	fmt.Println("refreshToken", refreshToken)
	if grantType != "refresh_token" {
		utils.ToJSON(&GenericResponse{Status: 400, Message: "Invalid grant type."}, w)
		return
	}

	connection, err := db.NewConnection()
	if err != nil {
		log.Fatalf("err:", err)
	}

	// リフレッシュトークンが失効されていないかチェック。
	disabled, _ := disabledRefreshTokenRepository.Exist(connection, refreshToken)
	if !disabled {
		utils.ToJSON(&GenericResponse{Status: 400, Message: "You are using disabled RefreshToken."}, w)
		return
	}

	// 正しいリフレッシュトークンかチェック。
	claims, err := utils.ValidateRefreshToken(refreshToken)
	if err != nil {
		utils.ToJSON(&GenericResponse{Status: 400, Message: "Authentication failed. Invalid token"}, w)
		return
	}

	// 正しいuserIdかチェックする。
	user, err := userRepositopry.FetchByUserId(connection, claims.UserID)
	if err != nil && user.Id != claims.UserID {
		utils.ToJSON(&GenericResponse{Status: 400, Message: "Invalid User."}, w)
		return
	}

	// 使用したリフレッシュトークンを失効
	disabledRefreshTokenRepository.AddDisabledRefreshToken(connection, refreshToken)
	// 新しいリフレッシュトークンを作成
	newRefreshToken, err := utils.GenerateRefreshToken(user.Id, user.TokenHash)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		utils.ToJSON(&GenericResponse{Status: http.StatusInternalServerError, Message: "Cannot Create RefreshToken"}, w)
		return
	}

	// 新しいアクセストークンを作成
	newAccessToken, err := utils.GenerateAccessToken(user.Id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		utils.ToJSON(&GenericResponse{Status: http.StatusInternalServerError, Message: "Cannot Create RefreshToken"}, w)
		return
	}
	// 作成した新しいトークンを返す
	data := &GenericResponse{
		Status:  http.StatusOK,
		Message: "Create New RefreshToken",
		Data:    &NewTokenResponse{AccessToken: newAccessToken, RefreshToken: newRefreshToken},
	}
	utils.ToJSON(data, w)
}
