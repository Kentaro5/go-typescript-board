package seeders

import (
	"fmt"
	"time"
)

type SexMst struct {
	code       int
	name       string
	created_at string
	updated_at string
}

func CustomerSeed() {
	// あくまでフォーマットで使用するため、適当な日付を指定する。
	dateTimeFormat := "2006-01-02 15:04:05"
	sexLists := []SexMst{
		SexMst{
			code:       0,
			name:       "不明",
			created_at: time.Now().Format(dateTimeFormat),
			updated_at: time.Now().Format(dateTimeFormat),
		},
		SexMst{
			code:       1,
			name:       "男性",
			created_at: time.Now().Format(dateTimeFormat),
			updated_at: time.Now().Format(dateTimeFormat),
		},
		SexMst{
			code:       2,
			name:       "女性",
			created_at: time.Now().Format(dateTimeFormat),
			updated_at: time.Now().Format(dateTimeFormat),
		},
		SexMst{
			code:       9,
			name:       "その他",
			created_at: time.Now().Format(dateTimeFormat),
			updated_at: time.Now().Format(dateTimeFormat),
		},
	}
	for index, sexMasterData := range sexLists {
		fmt.Println(index, sexMasterData.code)
		fmt.Println(index, sexMasterData.name)
	}
	//
	// for _, sexMasterData := range sexLists {
	// 	// prepare the statement
	// 	stmt, _ := s.db.Prepare(`INSERT INTO sex_mst(code, name, created_at, updated_at) VALUES (?,?,?,?)`)
	// 	// execute query
	// 	_, err := stmt.Exec(sexMasterData.code,
	// 		sexMasterData.name,
	// 		sexMasterData.created_at,
	// 		sexMasterData.updated_at,
	// 	)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }
}
