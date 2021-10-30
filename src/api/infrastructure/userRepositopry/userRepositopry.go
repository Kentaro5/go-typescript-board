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
	userResult := "user.id, user.name, user.email, user.sex_code, user.pref_code, user.city_code, user.ward_code, user.created_at, "
	sexMstResult := "sex_mst.code, sex_mst.name, "
	prefMstResult := "pref_mst.pref_code, pref_mst.pref, "
	cityMstResult := "city_mst.city_code, city_mst.city, "
	wardMstResult := "ward_mst.ward_code, ward_mst.ward "
	expectedResult := userResult + sexMstResult + prefMstResult + cityMstResult + wardMstResult
	var sql string = "SELECT DISTINCT " + expectedResult + "FROM " + userTableName + " as user "
	sql = sql + "JOIN sex_mst "
	sql = sql + "ON user.sex_code = sex_mst.code "
	sql = sql + "JOIN pref_mst "
	sql = sql + "ON user.pref_code = pref_mst.pref_code "
	sql = sql + "JOIN city_mst "
	sql = sql + "ON user.city_code = city_mst.city_code "
	sql = sql + "JOIN ward_mst "
	sql = sql + "ON user.ward_code = ward_mst.ward_code "
	sql = sql + "WHERE user.id = ?"
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
	var userSex UserSex
	var userPrefecture UserPrefecture
	var userCity UserCity
	var userWard UserWard

	for rows.Next() {
		// Goではループの終了後に、必ずエラーチェックを行う。
		// 全ての行が処理されるまで、ループが継続されるとは限らないため。
		// &形式で指定する。
		//  Scan error on column index 0, name "id": destination not a pointerのエラーが出るため
		err := rows.Scan(
			&user.Id,
			&user.Name,
			&user.Email,
			&user.SexCode,
			&user.PrefCode,
			&user.CityCode,
			&user.WardCode,
			&user.CreatedAt,
			&userSex.Code,
			&userSex.Name,
			&userPrefecture.PrefCode,
			&userPrefecture.Name,
			&userCity.CityCode,
			&userCity.Name,
			&userWard.WardCode,
			&userWard.Name,
		)
		if err != nil {
			log.Fatal(err)
		}
		// 参考記事：https://stackoverflow.com/questions/45637808/golang-db-query-with-sql-join
		user.Sex = append(user.Sex, userSex)
		user.Prefecture = append(user.Prefecture, userPrefecture)
		user.City = append(user.City, userCity)
		user.Ward = append(user.Ward, userWard)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return &user, nil
}

func UpdateByUserId(db *sql.DB, userId int, data UpdateUser) error {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
		return err
	}
	var sql string = "UPDATE " + userTableName + " "
	sql = sql + "set name = ?,  email = ?, sex_code = ?, pref_code = ?, city_code = ?, ward_code = ? "
	sql = sql + "where id = ?"

	// トランザクション管理処理
	defer func() {
		if err != nil {
			log.Fatal(err)
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	_, err = tx.Exec(
		sql,
		data.Name,
		data.Email,
		data.SexCode,
		data.PrefCode,
		data.CityCode,
		data.WardCode,
		userId,
	)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return err
}
