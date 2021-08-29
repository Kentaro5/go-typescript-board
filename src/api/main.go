package main

import (
	"api/Controllers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	//get := router.Methods(http.MethodGet).Subrouter()
	router.HandleFunc("/signUp", Controllers.SignUp).Methods("POST")

	// http.ListenAndServeで使用しているルーティングとポートを紐付けないと、動かない。
	log.Fatal(http.ListenAndServe(":8000", router))
}
