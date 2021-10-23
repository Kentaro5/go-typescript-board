package sexRepository

import (
	"database/sql"
	"log"
)

const sexMasterTableName string = "sex_mst"

func Fetch(db *sql.DB) (*Sexes, error) {

	var sql string = "SELECT id, name, code FROM " + sexMasterTableName
	stmt, err := db.Prepare(sql)
	if err != nil {
		return nil, err
	}

	// 遅延処理の実行
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	// 遅延処理の実行
	defer rows.Close()

	var sexes Sexes
	var sex Sex
	for rows.Next() {
		// Goではループの終了後に、必ずエラーチェックを行う。
		// 全ての行が処理されるまで、ループが継続されるとは限らないため。
		// &形式で指定する。
		//  Scan error on column index 0, name "id": destination not a pointerのエラーが出るため
		err := rows.Scan(
			&sex.Id,
			&sex.Name,
			&sex.Code,
		)
		if err != nil {
			log.Fatal(err)
		}
		sexes.SexLists = append(sexes.SexLists, sex)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return &sexes, nil
}
