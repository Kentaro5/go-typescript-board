package prefectureRepository

import (
	"database/sql"
	"log"
)

const prefectureMasterTableName string = "pref_mst"

func Fetch(db *sql.DB) (*Prefectures, error) {

	var sql string = "SELECT id, pref_code, pref FROM " + prefectureMasterTableName
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

	var prefectures Prefectures
	var prefecture Prefecture
	for rows.Next() {
		// Goではループの終了後に、必ずエラーチェックを行う。
		// 全ての行が処理されるまで、ループが継続されるとは限らないため。
		// &形式で指定する。
		//  Scan error on column index 0, name "id": destination not a pointerのエラーが出るため
		err := rows.Scan(
			&prefecture.Id,
			&prefecture.Code,
			&prefecture.Name,
		)
		if err != nil {
			log.Fatal(err)
		}
		prefectures.PrefectureLists = append(prefectures.PrefectureLists, prefecture)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return &prefectures, nil
}
