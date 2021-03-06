package Controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"api/db"
	"api/infrastructure/userRepositopry"
	"api/utils"
)

type Ping struct {
	Status int    `json:"status"`
	Result string `json:"result"`
}

// GenericResponse is the format of our response
type GenericResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type AuthResponse struct {
	RefreshToken string       `json:"refresh_token"`
	AccessToken  string       `json:"access_token"`
	User         ResponseUser `json:"user"`
}

type userFormData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ResponseUser struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func Login(w http.ResponseWriter, request *http.Request) {
	header := w.Header()
	header.Set("Content-Type", "application/json")
	header.Set("Access-Control-Allow-Origin", "http://localhost:3000")
	header.Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	// In case you don't have separate CORS middleware
	if request.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if request.Header.Get("Content-Type") != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
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

	var userData userFormData
	err = json.Unmarshal(body[:length], &userData)
	if err != nil {
		panic(err)
	}

	reqEmail := userData.Email
	reqPassword := userData.Password

	connection, err := db.NewConnection()
	if err != nil {
		utils.ToJSON(&GenericResponse{Status: 400, Message: "DB Connection Failed."}, w)
		return
	}

	user, err := userRepositopry.FetchByEmail(connection, reqEmail)
	if err != nil {
		utils.ToJSON(&GenericResponse{Status: 400, Message: "Login Failed."}, w)
		return
	}

	err = utils.CheckPassword(user.PasswordHash, reqPassword)
	if err != nil {
		utils.ToJSON(&GenericResponse{Status: 400, Message: "Password not corrected."}, w)
		return
	}

	accessToken, err := utils.GenerateAccessToken(user.Id)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		utils.ToJSON(&GenericResponse{Status: 400, Message: "Filed Generate Access Token."}, w)
		return
	}

	refreshToken, err := utils.GenerateRefreshToken(user.Id, user.TokenHash)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		utils.ToJSON(&GenericResponse{Status: 400, Message: "falied Generate Refresh Token."}, w)
		return
	}

	data := &GenericResponse{
		Status:  http.StatusOK,
		Message: "Successfully logged in",
		Data:    &AuthResponse{AccessToken: accessToken, RefreshToken: refreshToken, User: ResponseUser{Id: user.Id, Name: user.Name}},
	}
	utils.ToJSON(data, w)
}

func Sets(w http.ResponseWriter, request *http.Request) {

	// In case you don't have separate CORS middleware
	if request.Method == http.MethodOptions {
		header := w.Header()
		header.Set("Access-Control-Allow-Credentials", "true")
		header.Set("Access-Control-Allow-Headers", "Content-Type, withCredentials")
		header.Set("Access-Control-Allow-Origin", "http://localhost:3000")
		header.Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		w.WriteHeader(http.StatusOK)
		return
	}
	cookie := &http.Cookie{
		Name:     "test", // <- should be any unique key you want
		Value:    "test", // <- the token after encoded by SecureCookie
		Path:     "/",
		Secure:   true,
		HttpOnly: false,
		SameSite: http.SameSiteNoneMode,
		Domain:   "localhost",
	}
	http.SetCookie(w, cookie)

	ping := Ping{http.StatusOK, "ok"}
	res, _ := json.Marshal(ping)
	w.Write(res)
}

func PreflightSets(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	ping := Ping{http.StatusOK, "ok"}
	res, _ := json.Marshal(ping)
	w.Write(res)
}
