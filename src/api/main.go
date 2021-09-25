package main

import (
	"api/Controllers"
	"api/Middleware"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
)

func main() {
	amw := Middleware.AuthMiddleware{}
	amw.Initialize()
	router := mux.NewRouter()
	//get := router.Methods(http.MethodGet).Subrouter()
	router.HandleFunc("/signUp", Controllers.SignUp).Methods("POST")
	router.HandleFunc("/login", Controllers.Login).Methods("POST")
	router.HandleFunc("/login", Controllers.Login).Methods("OPTIONS")

	optionsRouter := router.Methods(http.MethodOptions).Subrouter()
	optionsRouter.HandleFunc("/", Controllers.Root)

	getRouter := router.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", Controllers.Root)
	getRouter.Use(amw.ValidateAccessToken)

	//cors optionsGoes Below
	c := cors.New(cors.Options{
		AllowedOrigins:     []string{"http://localhost:3000"}, // All origins
		AllowedMethods:     []string{"GET", "POST", "OPTIONS"},
		OptionsPassthrough: true,
	})
	// handler := cors.Default().Handler(router)
	handler := c.Handler(router)

	// http.ListenAndServeで使用しているルーティングとポートを紐付けないと、動かない。
	log.Fatal(http.ListenAndServe(":8000", handler))
}
