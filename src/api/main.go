package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"api/Domain/Entity/userEntity"
	"api/Domain/ValueObject/tokenValueObject"
	"api/db"
	"api/infrastructure/userRepositopry"
)

func main() {
	router := mux.NewRouter()
	//get := router.Methods(http.MethodGet).Subrouter()
	router.HandleFunc("/signUp", HomeHandler).Methods("POST")

	// http.ListenAndServeで使用しているルーティングとポートを紐付けないと、動かない。
	log.Fatal(http.ListenAndServe(":8000", router))
}

func HomeHandler(w http.ResponseWriter, request *http.Request) {
	connection, err := db.NewConnection()
	if err != nil {
		log.Fatalf("err:", err)
	}

	if err := request.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	userData := userEntity.User{}
	userData.Name = request.FormValue("Name")
	userData.Email = request.FormValue("Email")
	hashedPassword, err := userEntity.HashPassword(request.FormValue("Password"))
	if err != nil {
		fmt.Println("cannotHashPassword")
	}
	dateTimeFormat := "2006-01-02 15:04:05"
	dateTime := time.Now().Format(dateTimeFormat)

	// 各値をintにパース
	sexCode, err := strconv.ParseInt(request.FormValue("SexCode"), 10, 8)
	if err != nil {
		fmt.Println("cannot Convert SexCode")
	}
	prefCode, err := strconv.ParseInt(request.FormValue("PrefCode"), 10, 32)
	if err != nil {
		fmt.Println("cannot Convert PrefCode")
	}
	cityCode, err := strconv.ParseInt(request.FormValue("CityCode"), 10, 32)
	if err != nil {
		fmt.Println("cannot Convert CityCode")
	}
	wardCode, err := strconv.ParseInt(request.FormValue("WardCode"), 10, 32)
	if err != nil {
		fmt.Println("cannot Convert WardCode")
	}

	userData.PasswordHash = hashedPassword
	userData.TokenHash = tokenValueObject.GenerateRandomString(15)
	userData.SexCode = uint8(sexCode)
	userData.PrefCode = uint32(prefCode)
	userData.CityCode = uint32(cityCode)
	userData.WardCode = uint32(wardCode)
	userData.CreatedAt = dateTime
	userData.UpdatedAt = dateTime

	userRepositopry.Create(connection, userData)
	io.WriteString(w, "kitty kitty kitty")
}
