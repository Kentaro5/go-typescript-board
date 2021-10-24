package wardRepository

import (
	"database/sql"
	"log"
)

const wardMasterTableName string = "ward_mst"

func FetchByCityCode(db *sql.DB, prefCode int) (*Wards, error) {

	var sql string = "SELECT id, city_code, ward_code, ward FROM " + wardMasterTableName + " "
	sql = sql + "where city_code = ?"

	stmt, err := db.Prepare(sql)
	if err != nil {
		return nil, err
	}

	// 遅延処理の実行
	defer stmt.Close()
	rows, err := stmt.Query(prefCode)
	if err != nil {
		return nil, err
	}
	// 遅延処理の実行
	defer rows.Close()

	var wards Wards
	var ward Ward
	for rows.Next() {
		// Goではループの終了後に、必ずエラーチェックを行う。
		// 全ての行が処理されるまで、ループが継続されるとは限らないため。
		// &形式で指定する。
		//  Scan error on column index 0, name "id": destination not a pointerのエラーが出るため
		err := rows.Scan(
			&ward.Id,
			&ward.CityCode,
			&ward.Code,
			&ward.Name,
		)
		if err != nil {
			log.Fatal(err)
		}
		wards.WardLists = append(wards.WardLists, ward)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return &wards, nil
}
