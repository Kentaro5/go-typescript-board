package Controllers

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strconv"

	"api/db"
	"api/infrastructure/cityRepository"
	"api/utils"
	"github.com/gorilla/mux"
)

type citiesResponse struct {
	Cities []cityRepository.City `json:"cities"`
}

func GetCityLists(w http.ResponseWriter, request *http.Request) {
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
	prefCode, err := strconv.Atoi(params["prefCode"])
	if err != nil {
		utils.ToJSON(&GenericResponse{Status: 400, Message: "invalid parameters"}, w)
		return
	}

	cities, err := cityRepository.FetchByPrefCode(connection, prefCode)
	if err != nil {
		utils.ToJSON(&GenericResponse{Status: 400, Message: "Invalid User."}, w)
		return
	}

	fmt.Println(reflect.TypeOf(cities))
	fmt.Println("cities", cities)

	data := &GenericResponse{
		Status:  http.StatusOK,
		Message: "Successfully logged in",
		Data: &citiesResponse{
			Cities: cities.CityLists,
		},
	}

	utils.ToJSON(data, w)
}
