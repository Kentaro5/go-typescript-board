package userRepositopry

import (
	"api/Domain/Entity/userEntity"
	"database/sql"
	"log"
)

const userTableName string = "user_tbl"

func Create(db *sql.DB, user userEntity.User) {
	columns := "(name, email, password, token_hash, sex_code, pref_code, city_code, ward_code, created_at, updated_at)"

	sqlQuery, _ := db.Prepare("INSERT INTO " + userTableName + " " + columns + " VALUES (?,?,?,?,?,?,?,?,?,?)")
	// execute query
	_, err := sqlQuery.Exec(
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
	if err != nil {
		panic(err)
	}
}

func FetchByEmail(db *sql.DB, email string) (*User, error) {
	var sql string = "SELECT * FROM " + userTableName + " WHERE email = ?"
	stmt, err := db.Prepare(sql)
	if err != nil {
		return nil, err
	}
	// 遅延処理の実行
	defer stmt.Close()
	rows, err := stmt.Query(email)
	if err != nil {
		return nil, err
	}
	// 遅延処理の実行
	defer rows.Close()

	var user User
	for rows.Next() {
		// Goではループの終了後に、必ずエラーチェックを行う。
		// 全ての行が処理されるまで、ループが継続されるとは限らないため。
		// &形式で指定する。
		//  Scan error on column index 0, name "id": destination not a pointerのエラーが出るため
		err := rows.Scan(
			&user.Id,
			&user.Name,
			&user.Email,
			&user.PasswordHash,
			&user.TokenHash,
			&user.SexCode,
			&user.PrefCode,
			&user.CityCode,
			&user.WardCode,
			&user.RememberToken,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.DeletedAt,
		)
		if err != nil {
			log.Fatal(err)
		}
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return &user, nil
}

func FetchByUserId(db *sql.DB, userId int) (*User, error) {
	var sql string = "SELECT * FROM " + userTableName + " WHERE id = ?"
	stmt, err := db.Prepare(sql)
	if err != nil {
		return nil, err
	}
	// 遅延処理の実行
	defer stmt.Close()
	rows, err := stmt.Query(userId)
	if err != nil {
		return nil, err
	}
	// 遅延処理の実行
	defer rows.Close()

	var user User
	for rows.Next() {
		// Goではループの終了後に、必ずエラーチェックを行う。
		// 全ての行が処理されるまで、ループが継続されるとは限らないため。
		// &形式で指定する。
		//  Scan error on column index 0, name "id": destination not a pointerのエラーが出るため
		err := rows.Scan(
			&user.Id,
			&user.Name,
			&user.Email,
			&user.PasswordHash,
			&user.TokenHash,
			&user.SexCode,
			&user.PrefCode,
			&user.CityCode,
			&user.WardCode,
			&user.RememberToken,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.DeletedAt,
		)
		if err != nil {
			log.Fatal(err)
		}
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return &user, nil
}
