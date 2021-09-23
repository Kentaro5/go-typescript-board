package main

import (
	"api/Controllers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	//get := router.Methods(http.MethodGet).Subrouter()
	router.HandleFunc("/signUp", Controllers.SignUp).Methods("POST")
	router.HandleFunc("/login", Controllers.Login).Methods("POST")
	router.HandleFunc("/login", Controllers.Login).Methods("OPTIONS")
	router.HandleFunc("/", Controllers.Sets).Methods("POST")
	router.HandleFunc("/", Controllers.Sets).Methods("OPTIONS")

	//cors optionsGoes Below
	c := cors.New(cors.Options{
		AllowedOrigins:     []string{"http://localhost:3000"}, // All origins
		AllowCredentials:   true,                              // Cookieを共有できるようにセットしておく。
		AllowedMethods:     []string{"GET", "POST", "OPTIONS"},
		OptionsPassthrough: true,
	})
	// handler := cors.Default().Handler(router)
	handler := c.Handler(router)

	// http.ListenAndServeで使用しているルーティングとポートを紐付けないと、動かない。
	log.Fatal(http.ListenAndServe(":8000", handler))
}
