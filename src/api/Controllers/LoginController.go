package Controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"api/Domain/ValueObject/tokenValueObject"
	"api/db"
	"api/infrastructure/userRepositopry"
	"api/utils"
	"golang.org/x/crypto/bcrypt"
)

type Ping struct {
	Status int    `json:"status"`
	Result string `json:"result"`
}

// GenericResponse is the format of our response
type GenericResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type AuthResponse struct {
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
	Username     string `json:"username"`
}

type userFormData struct {
	Email    string
	Password string
}

func Login(w http.ResponseWriter, request *http.Request) {
	fmt.Println("Login:")
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		panic(err)
	}

	var userData userFormData
	err = json.Unmarshal([]byte(body), &userData)
	if err != nil {
		panic(err)
	}
	fmt.Println(userData.Email)

	reqEmail := userData.Email
	reqPassword := userData.Password

	fmt.Println("NewConnection:" + reqEmail)
	fmt.Println("NewConnection:" + reqPassword)

	connection, err := db.NewConnection()
	if err != nil {
		log.Fatalf("err:", err)
	}
	fmt.Println("NewConnection:")

	user, err := userRepositopry.FetchByEmail(connection, reqEmail)
	if err != nil {
		//errMsg := err.Error()
		//data.ToJSON(&GenericResponse{Status: false, Message: "Unable to retrieve user from database.Please try again later"}, w)
		return
	}
	fmt.Println("userRepositopry:")

	fmt.Println(&user)
	fmt.Println(reqPassword)

	result := checkPassword(user.PasswordHash, reqPassword)
	fmt.Println(result)
	if !result {
		return
	}
	fmt.Println("checkPassword:")

	userID := strconv.Itoa(user.Id)
	accessToken, err := tokenValueObject.GenerateAccessToken(userID)
	if err != nil {
		fmt.Println(accessToken)
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		// data.ToJSON(&GenericError{Error: err.Error()}, w)
		fmt.Println("cannot Convert SexCode")
		//utils.ToJSON(&GenericResponse{Status: false, Message: "Unable to login the user. Please try again later"}, w)
		return
	}
	fmt.Println("SexCode:")

	refreshToken, err := tokenValueObject.GenerateRefreshToken(userID, user.TokenHash)
	if err != nil {
		fmt.Println(accessToken)
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("cannotStatusInternalServerError")
		// data.ToJSON(&GenericError{Error: err.Error()}, w)
		//utils.ToJSON(&GenericResponse{Status: false, Message: "Unable to login the user. Please try again later"}, w)
		return
	}
	fmt.Println("AccessToken:")

	w.WriteHeader(http.StatusOK)
	// data.ToJSON(&AuthResponse{AccessToken: accessToken, RefreshToken: refreshToken, Username: user.Username}, w)
	utils.ToJSON(&GenericResponse{
		Status:  true,
		Message: "Successfully logged in",
		Data:    &AuthResponse{AccessToken: accessToken, RefreshToken: refreshToken, Username: user.Name},
	}, w)
}

func checkPassword(password string, requestPassword string) bool {
	// テキストのパスワードと、ハッシュ化されたパスワードを比較する
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(requestPassword))
	if err != nil {
		return false
	}

	return true
}

func Sets(w http.ResponseWriter, request *http.Request) {

	// In case you don't have separate CORS middleware
	if request.Method == http.MethodOptions {
		fmt.Println(http.SameSiteNoneMode)
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
	fmt.Println(http.SameSiteNoneMode)

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
