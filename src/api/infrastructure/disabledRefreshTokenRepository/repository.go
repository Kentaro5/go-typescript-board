package disabledRefreshTokenRepository

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

const disabledRefreshTokenTableName string = "disabled_refresh_token_tbl"

func AddDisabledRefreshToken(db *sql.DB, refreshToken string) (bool, error) {
	dateTimeFormat := "2006-01-02 15:04:05"
	dateTime := time.Now().Format(dateTimeFormat)
	columns := "(refresh_token, created_at, updated_at)"

	sqlQuery, _ := db.Prepare("INSERT INTO " + disabledRefreshTokenTableName + " " + columns + " VALUES (?,?,?)")
	// execute query
	_, err := sqlQuery.Exec(
		refreshToken,
		dateTime,
		dateTime,
	)
	if err != nil {
		panic(err)
	}

	return true, nil
}

func Exist(db *sql.DB, refreshToken string) (bool, error) {
	var sql string = "SELECT count(id) FROM " + disabledRefreshTokenTableName + " WHERE refresh_token = ?"
	stmt, err := db.Prepare(sql)
	if err != nil {
		return false, err
	}
	// 遅延処理の実行
	defer stmt.Close()
	rows, err := stmt.Query(refreshToken)
	if err != nil {
		return false, err
	}
	// 遅延処理の実行
	defer rows.Close()

	var count int
	for rows.Next() {
		if err := rows.Scan(&count); err != nil {
			log.Fatal(err)
		}
	}

	fmt.Printf("Number of rows are %s\n", count)

	if count < 0 {
		return false, nil
	}

	return true, nil
}
