package userRepositopry

import (
	"api/Domain/Entity/userEntity"
	"database/sql"
	"fmt"
)

const userDbName string = "user_tbl"

func Create(db *sql.DB, user userEntity.User) {
	columns := "(name, email, password, token_hash, sex_code, pref_code, city_code, ward_code, created_at, updated_at)"

	fmt.Println("test")

	stmt, _ := db.Prepare("INSERT INTO " + userDbName + " " + columns + " VALUES (?,?,?,?,?,?,?,?,?,?)")
	// execute query
	_, err := stmt.Exec(
		user.Name,
		user.Email,
		user.PasswordHash,
		user.TokenHash,
		user.SexCode,
		user.PrefCode,
		user.CityCode,
		user.WardCode,
		user.CreatedAt,
		user.UpdatedAt,
	)
	fmt.Println("test222")
	if err != nil {
		panic(err)
	}
}
