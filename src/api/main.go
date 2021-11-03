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
	//認証が必要ない場合は、以下は、routerで行う
	router.HandleFunc("/signUp", Controllers.SignUp).Methods("POST")
	router.HandleFunc("/login", Controllers.Login).Methods("POST")
	router.HandleFunc("/login", Controllers.Login).Methods("OPTIONS")
	router.HandleFunc("/get/sex", Controllers.GetSexLists).Methods("GET")
	router.HandleFunc("/get/prefecture", Controllers.GetPrefectureLists).Methods("GET")
	router.HandleFunc("/get/city/{prefCode}", Controllers.GetCityLists).Methods("GET")
	router.HandleFunc("/get/ward/{cityCode}", Controllers.GetWardLists).Methods("GET")
	router.HandleFunc("/refresh-token", Controllers.CreateAccessTokenByRefreshToken).Methods("POST")

	optionsRouter := router.Methods(http.MethodOptions).Subrouter()
	optionsRouter.HandleFunc("/", Controllers.Root)
	optionsRouter.HandleFunc("/signUp", Controllers.OptionUser)
	optionsRouter.HandleFunc("/user/{userId}", Controllers.OptionUser)
	optionsRouter.HandleFunc("/user/{userId}/changePassword/", Controllers.OptionUser)
	optionsRouter.HandleFunc("/sex", Controllers.GetSexLists)
	optionsRouter.HandleFunc("/prefecture", Controllers.GetPrefectureLists)
	optionsRouter.HandleFunc("/city/{prefCode}", Controllers.GetCityLists)
	optionsRouter.HandleFunc("/ward/{cityCode}", Controllers.GetWardLists)

	getRouter := router.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", Controllers.Root)
	getRouter.HandleFunc("/user/{userId}", Controllers.GetUser)
	getRouter.HandleFunc("/sex", Controllers.GetSexLists)
	getRouter.HandleFunc("/prefecture", Controllers.GetPrefectureLists)
	getRouter.HandleFunc("/city/{prefCode}", Controllers.GetCityLists)
	getRouter.HandleFunc("/ward/{cityCode}", Controllers.GetWardLists)
	getRouter.Use(amw.ValidateAccessToken)

	patchRouter := router.Methods(http.MethodPatch).Subrouter()
	patchRouter.HandleFunc("/user/{userId}", Controllers.UpdateUser)
	patchRouter.HandleFunc("/user/{userId}/changePassword/", Controllers.ChangeUserPassword)
	patchRouter.Use(amw.ValidateAccessToken)

	//cors optionsGoes Below
	c := cors.New(cors.Options{
		AllowedOrigins:     []string{"http://localhost:3000"}, // All origins
		AllowedMethods:     []string{"GET", "POST", "OPTIONS", "PATCH"},
		OptionsPassthrough: true,
	})
	// handler := cors.Default().Handler(router)
	handler := c.Handler(router)

	// http.ListenAndServeで使用しているルーティングとポートを紐付けないと、動かない。
	log.Fatal(http.ListenAndServe(":8000", handler))
}
