package main

import (
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"os"

	"api/db"
	"api/infrastructure/userRepositopry"
)

func main() {
	router := mux.NewRouter()
	//get := router.Methods(http.MethodGet).Subrouter()
	post := router.Methods(http.MethodPost.Subrouter())
	post.HandleFunc("/", HomeHandler)

	// http.ListenAndServeで使用しているルーティングとポートを紐付けないと、動かない。
	log.Fatal(http.ListenAndServe(":8000", router))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	connection, err := db.NewConnection()
	if err != nil {
		log.Fatalf("err:", err)
	}

	os.Exit(1)
	userData := userRepositopry.User{
		Name:         "test",
		Email:        "test@exanple.com",
		PasswordHash: "test",
		SexCode:      0,
		PrefCode:     10006,
		CityCode:     11002,
		WardCode:     11011,
		CreatedAt:    "2021-09-09 12:00:00",
		UpdatedAt:    "2021-09-09 12:00:00",
	}

	userRepositopry.Create(connection, userData)
	io.WriteString(w, "kitty kitty kitty")
}
