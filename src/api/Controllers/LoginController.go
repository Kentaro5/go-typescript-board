package Controllers

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"api/Domain/ValueObject/tokenValueObject"
	"api/db"
	"api/infrastructure/userRepositopry"
	"golang.org/x/crypto/bcrypt"
	"utils"
)

func Login(w http.ResponseWriter, request *http.Request) {
	reqEmail := request.FormValue("email")
	reqPassword := request.FormValue("password")

	connection, err := db.NewConnection()
	if err != nil {
		log.Fatalf("err:", err)
	}

	user, err := userRepositopry.FetchByEmail(connection, reqEmail)
	if err != nil {
		//errMsg := err.Error()
		//data.ToJSON(&GenericResponse{Status: false, Message: "Unable to retrieve user from database.Please try again later"}, w)
		return
	}

	result := checkPassword(user.PasswordHash, reqPassword)
	if !result {
		return
	}

	userID := strconv.Itoa(user.Id)
	accessToken, err := tokenValueObject.GenerateAccessToken(userID)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		// data.ToJSON(&GenericError{Error: err.Error()}, w)
		utils.ToJSON(&GenericResponse{Status: false, Message: "Unable to login the user. Please try again later"}, w)
		return
	}

	fmt.Println(accessToken)
}

func checkPassword(password string, requestPassword string) bool {
	// テキストのパスワードと、ハッシュ化されたパスワードを比較する
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(requestPassword))
	if err != nil {
		return false
	}

	return true
}
