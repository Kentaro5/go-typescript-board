package main

import (
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	getRouter := router.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", HomeHandler)

	// http.ListenAndServeで使用しているルーティングとポートを紐付けないと、動かない。
	log.Fatal(http.ListenAndServe(":8000", router))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "kitty kitty kitty")
}
