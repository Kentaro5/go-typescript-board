package Controllers

import (
	"fmt"
	"log"
	"net/http"
	"reflect"

	"api/db"
	"api/infrastructure/sexRepository"
	"api/utils"
)

type sexListsResponse struct {
	Sexes []sexRepository.Sex `json:"sexes"`
}

func GetSexLists(w http.ResponseWriter, request *http.Request) {
	header := w.Header()
	header.Set("Content-Type", "application/json")
	header.Set("Access-Control-Allow-Origin", "http://localhost:3000")
	header.Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	header.Set("Access-Control-Allow-Headers", "Authorization")

	// In case you don't have separate CORS middleware
	if request.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	connection, err := db.NewConnection()
	if err != nil {
		log.Fatalf("err:", err)
	}

	sexes, err := sexRepository.Fetch(connection)
	if err != nil {
		utils.ToJSON(&GenericResponse{Status: 400, Message: "Invalid User."}, w)
		return
	}
	fmt.Println("sexes.SexLists", sexes.SexLists)
	fmt.Println("sexes.SexLists", reflect.TypeOf(sexes.SexLists))
	data := &GenericResponse{
		Status:  http.StatusOK,
		Message: "Successfully logged in",
		Data: &sexListsResponse{
			Sexes: sexes.SexLists,
		},
	}
	fmt.Println("sexes.SexLists", data)
	utils.ToJSON(data, w)
}
