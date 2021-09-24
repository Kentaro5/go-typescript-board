package Controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"api/Domain/Entity/userEntity"
	"api/db"
	"api/infrastructure/userRepositopry"
	"api/utils"
)

func SignUp(w http.ResponseWriter, request *http.Request) {
	if err := request.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	userData := userEntity.User{}
	userData.Name = request.FormValue("name")
	userData.Email = request.FormValue("email")
	hashedPassword, err := userEntity.HashPassword(request.FormValue("password"))
	if err != nil {
		fmt.Println("cannotHashPassword")
	}
	dateTimeFormat := "2006-01-02 15:04:05"
	dateTime := time.Now().Format(dateTimeFormat)

	// 各値をintにパース
	sexCode, err := strconv.ParseInt(request.FormValue("sex_code"), 10, 8)
	if err != nil {
		fmt.Println("cannot Convert SexCode")
	}
	prefCode, err := strconv.ParseInt(request.FormValue("pref_code"), 10, 32)
	if err != nil {
		fmt.Println("cannot Convert PrefCode")
	}
	cityCode, err := strconv.ParseInt(request.FormValue("city_code"), 10, 32)
	if err != nil {
		fmt.Println("cannot Convert CityCode")
	}
	wardCode, err := strconv.ParseInt(request.FormValue("ward_code"), 10, 32)
	if err != nil {
		fmt.Println("cannot Convert WardCode")
	}

	userData.PasswordHash = hashedPassword
	userData.TokenHash = utils.GenerateRandomString(15)
	userData.SexCode = uint8(sexCode)
	userData.PrefCode = uint32(prefCode)
	userData.CityCode = uint32(cityCode)
	userData.WardCode = uint32(wardCode)
	userData.CreatedAt = dateTime
	userData.UpdatedAt = dateTime

	connection, err := db.NewConnection()
	if err != nil {
		log.Fatalf("err:", err)
	}
	userRepositopry.Create(connection, userData)

	http.Redirect(w, request, "http://localhost:3000/login", http.StatusSeeOther)
}
