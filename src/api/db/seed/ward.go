package seed

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"path/filepath"
	"strconv"
	"time"
)

func (s Seed) WardMstSeed() {
	// 相対パスから絶対パスを取得。
	filePath, err := filepath.Abs("../assets/files/ward/ward.xlsx")
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
	var cityCode int
	for i := 1; i < len(rows); i++ {
		// 都道府県が存在し、かつ市区町村の箇所が空の場合は、都道府県が切り替わるタイミングなので、その時に都道府県コードをセットする。
		if rows[i][ward_Index] == "" {
			// intergerに変換
			cityCode, _ = strconv.Atoi(rows[i][code_Index])
		} else {
			stmt, prepareErrors := s.db.Prepare(`INSERT INTO ward_mst(city_code, ward_code, ward, created_at, updated_at) VALUES (?,?,?,?,?)`)
			if prepareErrors != nil {
				panic(prepareErrors)
			}
			// execute query
			_, execError := stmt.Exec(
				cityCode,
				rows[i][ward_code_Index],
				rows[i][ward_Index],
				time.Now().Format(dateTimeFormat),
				time.Now().Format(dateTimeFormat),
			)
			if execError != nil {
				panic(execError)
			}
		}
	}
}
