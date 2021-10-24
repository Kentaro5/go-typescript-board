package Controllers

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strconv"

	"api/db"
	"api/infrastructure/wardRepository"
	"api/utils"
	"github.com/gorilla/mux"
)

type wardsResponse struct {
	Wards []wardRepository.Ward `json:"wards"`
}

func GetWardLists(w http.ResponseWriter, request *http.Request) {
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

	params := mux.Vars(request)
	cityCode, err := strconv.Atoi(params["cityCode"])
	if err != nil {
		utils.ToJSON(&GenericResponse{Status: 400, Message: "invalid parameters"}, w)
		return
	}

	wards, err := wardRepository.FetchByCityCode(connection, cityCode)
	if err != nil {
		utils.ToJSON(&GenericResponse{Status: 400, Message: "Invalid User."}, w)
		return
	}

	fmt.Println(reflect.TypeOf(wards))
	fmt.Println("wards", wards)

	data := &GenericResponse{
		Status:  http.StatusOK,
		Message: "Successfully logged in",
		Data: &wardsResponse{
			Wards: wards.WardLists,
		},
	}

	utils.ToJSON(data, w)
}
