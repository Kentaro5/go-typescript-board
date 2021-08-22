package seed

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"path/filepath"
	"time"
)

func (s Seed) PrefMstSeed() {
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

	for _, row := range rows {
		if row[prefecture_Index] != "" && row[city_Index] == "" {
			stmt, errors := s.db.Prepare(`INSERT INTO pref_mst(pref_code, pref, created_at, updated_at) VALUES (?,?,?,?)`)
			// execute query
			_, err := stmt.Exec(
				row[code_Index], // code_idexが整数ではなく、文字列で来るのにGoでSQLインサートする際に暗黙的に、文字列から整数に変換されている。
				row[prefecture_Index],
				time.Now().Format(dateTimeFormat),
				time.Now().Format(dateTimeFormat),
			)
			if err != nil {
				panic(err)
			}
		}
	}
}
