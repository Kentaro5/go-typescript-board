package Controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"api/Domain/UseCase/PasswordUseCase"
	"api/Domain/UseCase/TokenUseCase"
	"api/db"
	"api/infrastructure/userRepositopry"
	"api/utils"
	"github.com/gorilla/mux"
)

type response struct {
	Id         int     `json:"user_id"`
	Name       string  `json:"user_name"`
	Email      string  `json:"email"`
	SexCode    uint8   `json:"sex_code"`
	Sex        string  `json:"sex"`
	PrefCode   uint32  `json:"pref_code"`
	Prefecture string  `json:"prefecture"`
	CityCode   uint32  `json:"city_code"`
	City       string  `json:"city"`
	WardCode   *uint32 `json:"ward_code"`
	Ward       *string `json:"ward"`
	CreatedAt  string  `json:"created_at"`
}

func OptionUser(w http.ResponseWriter, request *http.Request) {
	header := w.Header()
	header.Set("Content-Type", "application/json")
	header.Set("Access-Control-Allow-Origin", "http://localhost:3000")
	header.Set("Access-Control-Allow-Methods", "GET, PATCH, OPTIONS")
	header.Set("Access-Control-Allow-Headers", "Authorization, Content-Type") // PATCHメソッドの場合は、content-typeも付与する。

	// In case you don't have separate CORS middleware
	if request.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
}

func GetUser(w http.ResponseWriter, request *http.Request) {
	header := w.Header()
	header.Set("Content-Type", "application/json")
	header.Set("Access-Control-Allow-Origin", "http://localhost:3000")
	header.Set("Access-Control-Allow-Methods", "GET")
	header.Set("Access-Control-Allow-Headers", "Authorization")

	connection, err := db.NewConnection()
	if err != nil {
		log.Fatalf("err:", err)
	}

	params := mux.Vars(request)
	userId, err := strconv.Atoi(params["userId"])
	if err != nil {
		utils.ToJSON(&GenericResponse{Status: 400, Message: "invalid parameters"}, w)
		return
	}

	user, err := userRepositopry.FetchByUserId(connection, userId)
	if err != nil && user.Id != userId {
		utils.ToJSON(&GenericResponse{Status: 400, Message: "Invalid User."}, w)
		return
	}

	//utils.ToJSON(&AuthResponse{AccessToken: accessToken, RefreshToken: refreshToken, Username: user.Username}, w)
	data := &GenericResponse{
		Status:  http.StatusOK,
		Message: "Successfully logged in",
		Data: &response{
			Id:         user.Id,
			Name:       user.Name,
			Email:      user.Email,
			SexCode:    user.SexCode,
			Sex:        user.Sex[0].Name,
			PrefCode:   user.PrefCode,
			Prefecture: user.Prefecture[0].Name,
			CityCode:   user.CityCode,
			City:       user.City[0].Name,
			WardCode:   user.WardCode,
			Ward:       user.Ward[0].Name,
			CreatedAt:  user.CreatedAt,
		},
	}

	utils.ToJSON(data, w)
}

func UpdateUser(w http.ResponseWriter, request *http.Request) {
	header := w.Header()
	header.Set("Content-Type", "application/json")
	header.Set("Access-Control-Allow-Origin", "http://localhost:3000")
	header.Set("Access-Control-Allow-Methods", "PATCH")
	header.Set("Access-Control-Allow-Headers", "Authorization")

	// In case you don't have separate CORS middleware
	if request.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	// In case you don't have separate CORS middleware
	if request.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	params := mux.Vars(request)
	userId, err := strconv.Atoi(params["userId"])
	if err != nil {
		utils.ToJSON(&GenericResponse{Status: 400, Message: "invalid parameters"}, w)
		return
	}

	//To allocate slice for request body
	length, err := strconv.Atoi(request.Header.Get("Content-Length"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//Read body data to parse json
	body := make([]byte, length)
	length, err = request.Body.Read(body)
	if err != nil && err != io.EOF {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var updateData userRepositopry.UpdateUser
	err = json.Unmarshal(body[:length], &updateData)
	if err != nil {
		panic(err)
	}

	reqName := updateData.Name
	reqSexCode := updateData.SexCode
	reqEmail := updateData.Email
	reqPrefCode := updateData.PrefCode
	reqCityCode := updateData.CityCode
	reqWardCode := updateData.WardCode

	fmt.Println("reqName:", reqName)
	fmt.Println("reqSexCode:", reqSexCode)
	fmt.Println("reqEmail:", reqEmail)
	fmt.Println("reqPrefCode:", reqPrefCode)
	fmt.Println("reqCityCode:", reqCityCode)
	fmt.Println("reqWardCode:", reqWardCode)

	fmt.Println(userId)
	connection, err := db.NewConnection()
	if err != nil {
		log.Fatalf("err:", err)
	}
	fmt.Println("reqWardCode:", reqWardCode)
	err = userRepositopry.UpdateByUserId(connection, userId, updateData)
	fmt.Println("reqWardCode:", err)
	if err != nil {
		utils.ToJSON(&GenericResponse{Status: 400, Message: "Failed Update User."}, w)
		return
	}

	// user, err := userRepositopry.FetchByUserId(connection, userId)
	// if err != nil && user.Id != userId {
	// 	utils.ToJSON(&GenericResponse{Status: 400, Message: "Invalid User."}, w)
	// 	return
	// }

	//utils.ToJSON(&AuthResponse{AccessToken: accessToken, RefreshToken: refreshToken, Username: user.Username}, w)
	// data := &GenericResponse{
	// 	Status:  http.StatusOK,
	// 	Message: "Successfully logged in",
	// 	Data: &response{
	// 		Id:         user.Id,
	// 		Name:       user.Name,
	// 		Email:      user.Email,
	// 		SexCode:    user.SexCode,
	// 		Sex:        user.Sex[0].Name,
	// 		PrefCode:   user.PrefCode,
	// 		Prefecture: user.Prefecture[0].Name,
	// 		CityCode:   user.CityCode,
	// 		City:       user.City[0].Name,
	// 		WardCode:   user.WardCode,
	// 		Ward:       user.Ward[0].Name,
	// 		CreatedAt:  user.CreatedAt,
	// 	},
	// }
	//
	// utils.ToJSON(data, w)
}

func ChangeUserPassword(w http.ResponseWriter, request *http.Request) {
	// 送られてきたパスワードで、ユーザー情報を引き当てる。
	// 引き当てれなかったら、エラー
	// 引き当てたら正しいので、パスワード変更処理を行う
	// パスワードを再度ハッシュ化する。
	// パスワードのハッシュ化ができたら、今度は各トークンを新しくする。
	// 送られてきた
	header := w.Header()
	header.Set("Content-Type", "application/json")
	header.Set("Access-Control-Allow-Origin", "http://localhost:3000")
	header.Set("Access-Control-Allow-Methods", "PATCH")
	header.Set("Access-Control-Allow-Headers", "Authorization")

	params := mux.Vars(request)
	userId, err := strconv.Atoi(params["userId"])
	if err != nil {
		fmt.Println("test", userId)
		utils.ToJSON(&GenericResponse{Status: 400, Message: "invalid userId parameters"}, w)
		return
	}

	//To allocate slice for request body
	length, err := strconv.Atoi(request.Header.Get("Content-Length"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//Read body data to parse json
	body := make([]byte, length)
	length, err = request.Body.Read(body)
	if err != nil && err != io.EOF {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var updateData userRepositopry.UpdatePassword
	err = json.Unmarshal(body[:length], &updateData)
	if err != nil {
		panic(err)
	}

	oldPassword := updateData.OldPassword
	newPassword := updateData.NewPassword
	grantType := updateData.GrantType
	refreshToken := updateData.RefreshToken

	fmt.Println("oldPassword", oldPassword)
	fmt.Println("newPassword", newPassword)

	err = PasswordUseCase.ChangePassword(userId, oldPassword, newPassword)
	if err != nil {
		fmt.Println(err)
		utils.ToJSON(&GenericResponse{Status: 400, Message: "invalid parameters"}, w)
		return
	}

	fmt.Println("grantType", grantType)
	fmt.Println("refreshToken", refreshToken)
	if grantType != "refresh_token" {
		utils.ToJSON(&GenericResponse{Status: 400, Message: "Invalid grant type."}, w)
		return
	}

	refreshToken, err = TokenUseCase.RefreshToken(refreshToken)
	if err != nil {
		fmt.Println(err)
		utils.ToJSON(&GenericResponse{Status: 400, Message: "You are using disabled Token"}, w)
		return
	}

	if refreshToken == "" {
		fmt.Println(err)
		utils.ToJSON(&GenericResponse{Status: 400, Message: "invalid parameters"}, w)
		return
	}

	// 新しいアクセストークンを作成
	newAccessToken, err := utils.GenerateAccessToken(userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		utils.ToJSON(&GenericResponse{Status: http.StatusInternalServerError, Message: "Cannot Create RefreshToken"}, w)
		return
	}

	// 作成した新しいトークンを返す
	data := &GenericResponse{
		Status:  http.StatusOK,
		Message: "Create New RefreshToken",
		Data:    &NewTokenResponse{AccessToken: newAccessToken, RefreshToken: refreshToken},
	}
	utils.ToJSON(data, w)
}
