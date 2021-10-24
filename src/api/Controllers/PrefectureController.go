package Controllers

import (
	"fmt"
	"log"
	"net/http"
	"reflect"

	"api/db"
	"api/infrastructure/prefectureRepository"
	"api/utils"
)

type prefecturesResponse struct {
	Prefectures []prefectureRepository.Prefecture `json:"prefectures"`
}

func GetPrefectureLists(w http.ResponseWriter, request *http.Request) {
	fmt.Println("prefectures2222")
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
	fmt.Println("prefectures2222")
	prefectures, err := prefectureRepository.Fetch(connection)
	if err != nil {
		utils.ToJSON(&GenericResponse{Status: 400, Message: "Invalid User."}, w)
		return
	}

	fmt.Println("prefectures", prefectures)

	fmt.Println(reflect.TypeOf(prefectures))

	data := &GenericResponse{
		Status:  http.StatusOK,
		Message: "Successfully logged in",
		Data: &prefecturesResponse{
			Prefectures: prefectures.PrefectureLists,
		},
	}

	utils.ToJSON(data, w)
}
