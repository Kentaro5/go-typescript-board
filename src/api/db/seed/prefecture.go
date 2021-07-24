package seed

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"path/filepath"
	"time"
)

type PrefMst struct {
	pref_code  int
	pref       string
	created_at string
	updated_at string
}

const (
	CODE_INDEX       int = 0
	PREFECTURE_INDEX int = 1
	CITY_INDEX       int = 2
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
		if row[PREFECTURE_INDEX] != "" && row[CITY_INDEX] == "" {
			fmt.Print(row[PREFECTURE_INDEX], "\t")
			stmt, errors := s.db.Prepare(`INSERT INTO pref_mst(pref_code, pref, created_at, updated_at) VALUES (?,?,?,?)`)
			fmt.Println(stmt, errors)
			// execute query
			_, err := stmt.Exec(
				row[CODE_INDEX],
				row[PREFECTURE_INDEX],
				time.Now().Format(dateTimeFormat),
				time.Now().Format(dateTimeFormat),
			)
			if err != nil {
				panic(err)
			}
		}
	}

}
