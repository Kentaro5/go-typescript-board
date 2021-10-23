package Controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"api/db"
	"api/infrastructure/userRepositopry"
	"api/utils"
	"github.com/gorilla/mux"
)

type params struct {
	Id int
}

type response struct {
	Id         int     `json:"user_id"`
	Name       string  `json:"user_name"`
	Email      string  `json:"email"`
	SexCode    uint8   `json:"sex_code"`
	Sex        string  `json:"sex"`
	PrefCode   uint32  `json:"pref_code"`
	Prefecture string  `json:"prefecture"`
	CityCode   uint32  `json:"city_code"`
	City       string  `json:"city"`
	WardCode   *uint32 `json:"ward_code"`
	Ward       *string `json:"ward"`
	CreatedAt  string  `json:"created_at"`
}

func GetUser(w http.ResponseWriter, request *http.Request) {
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
	userId, err := strconv.Atoi(params["userId"])
	if err != nil {
		utils.ToJSON(&GenericResponse{Status: 400, Message: "invalid parameters"}, w)
		return
	}

	user, err := userRepositopry.FetchByUserId(connection, userId)
	if err != nil && user.Id != userId {
		utils.ToJSON(&GenericResponse{Status: 400, Message: "Invalid User."}, w)
		return
	}

	//utils.ToJSON(&AuthResponse{AccessToken: accessToken, RefreshToken: refreshToken, Username: user.Username}, w)
	data := &GenericResponse{
		Status:  http.StatusOK,
		Message: "Successfully logged in",
		Data: &response{
			Id:         user.Id,
			Name:       user.Name,
			Email:      user.Email,
			SexCode:    user.SexCode,
			Sex:        user.Sex[0].Name,
			PrefCode:   user.PrefCode,
			Prefecture: user.Prefecture[0].Name,
			CityCode:   user.CityCode,
			City:       user.City[0].Name,
			WardCode:   user.WardCode,
			Ward:       user.Ward[0].Name,
			CreatedAt:  user.CreatedAt,
		},
	}
	fmt.Println("ToJSON", data)
	utils.ToJSON(data, w)
}
