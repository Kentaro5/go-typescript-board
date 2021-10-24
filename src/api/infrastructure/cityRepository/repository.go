package cityRepository

import (
	"database/sql"
	"log"
)

const cityMasterTableName string = "city_mst"

func FetchByPrefCode(db *sql.DB, prefCode int) (*Cities, error) {

	var sql string = "SELECT id, pref_code, city_code, city FROM " + cityMasterTableName + " "
	sql = sql + "where pref_code = ?"

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

	var cities Cities
	var city City
	for rows.Next() {
		// Goではループの終了後に、必ずエラーチェックを行う。
		// 全ての行が処理されるまで、ループが継続されるとは限らないため。
		// &形式で指定する。
		//  Scan error on column index 0, name "id": destination not a pointerのエラーが出るため
		err := rows.Scan(
			&city.Id,
			&city.PrefCode,
			&city.Code,
			&city.Name,
		)
		if err != nil {
			log.Fatal(err)
		}
		cities.CityLists = append(cities.CityLists, city)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return &cities, nil
}
