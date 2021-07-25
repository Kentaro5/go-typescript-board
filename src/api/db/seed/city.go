package seed

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"path/filepath"
	"strconv"
	"time"
)

type CityMst struct {
	pref_code  int
	pref       string
	created_at string
	updated_at string
}

func (s Seed) CityMstSeed() {
	// 相対パスから絶対パスを取得。
	filePath, err := filepath.Abs("../assets/files/pref/pref.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}

	f, err := excelize.OpenFile(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}

	// あくまでフォーマットで使用するため、適当な日付を指定する。
	dateTimeFormat := "2006-01-02 15:04:05"
	var prefCode int

	for i := 1; i < len(rows); i++ {
		// 都道府県が存在し、かつ市区町村の箇所が空の場合は、都道府県が切り替わるタイミングなので、その時に都道府県コードをセットする。
		if rows[i][prefecture_Index] != "" && rows[i][city_Index] == "" {
			// intergerに変換
			prefCode, _ = strconv.Atoi(rows[i][code_Index])
		}

		if rows[i][prefecture_Index] != "" && rows[i][city_Index] != "" {
			fmt.Print(rows[i][prefecture_Index], "\t")
			stmt, errors := s.db.Prepare(`INSERT INTO city_mst(pref_code, city_code, city, created_at, updated_at) VALUES (?,?,?,?,?)`)
			fmt.Println(stmt, errors)
			// execute query
			_, err := stmt.Exec(
				prefCode,
				rows[i][code_Index],
				rows[i][city_Index],
				time.Now().Format(dateTimeFormat),
				time.Now().Format(dateTimeFormat),
			)
			fmt.Println(prefCode,
				rows[i][code_Index],
				rows[i][city_Index],
				time.Now().Format(dateTimeFormat),
				time.Now().Format(dateTimeFormat))
			if err != nil {
				panic(err)
			}
		}
	}
}
