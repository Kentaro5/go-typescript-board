package Controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"api/db"
	"api/infrastructure/userRepositopry"
	"api/utils"
)

type SignUpFormData struct {
	Name      string  `json:"name"`
	Email     string  `json:"email"`
	Password  string  `json:"password"`
	SexCode   uint8   `json:"sex_code"`
	PrefCode  uint32  `json:"pref_code"`
	CityCode  uint32  `json:"city_code"`
	WardCode  *uint32 `json:"ward_code"`
	CreatedAt *string
	UpdatedAt *string
}

func SignUp(w http.ResponseWriter, request *http.Request) {
	header := w.Header()
	header.Set("Content-Type", "application/json")
	header.Set("Access-Control-Allow-Origin", "http://localhost:3000")
	header.Set("Access-Control-Allow-Methods", "POST")

	//To allocate slice for request body
	length, err := strconv.Atoi(request.Header.Get("Content-Length"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var userData userRepositopry.SignUpUser
	dateTimeFormat := "2006-01-02 15:04:05"
	dateTime := time.Now().Format(dateTimeFormat)

	//Read body data to parse json
	body := make([]byte, length)
	length, err = request.Body.Read(body)
	if err != nil && err != io.EOF {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var requestData SignUpFormData
	err = json.Unmarshal(body[:length], &requestData)
	if err != nil {
		panic(err)
	}

	hashedPassword, err := utils.HashPassword(requestData.Password)
	if err != nil {
		log.Fatalf("err:", err)
	}
	fmt.Println("requestData.Password", requestData.Password)

	err = utils.CheckPassword(hashedPassword, requestData.Password)
	if err != nil {
		utils.ToJSON(&GenericResponse{Status: 400, Message: "Password not corrected."}, w)
		return
	}

	userData.Name = requestData.Name
	userData.Email = requestData.Email
	userData.PasswordHash = hashedPassword
	userData.TokenHash = utils.GenerateRandomString(15)
	userData.SexCode = requestData.SexCode
	userData.PrefCode = requestData.PrefCode
	userData.CityCode = requestData.CityCode
	userData.WardCode = requestData.WardCode
	userData.CreatedAt = dateTime
	userData.UpdatedAt = dateTime

	connection, err := db.NewConnection()
	if err != nil {
		log.Fatalf("err:", err)
	}
	userRepositopry.Create(connection, userData)

	utils.ToJSON(&GenericResponse{Status: http.StatusOK, Message: "SignUpOK"}, w)
	return
}
