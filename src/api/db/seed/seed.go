package seed

import (
	"database/sql"
	"log"
	"reflect"
)

type Seed struct {
	db *sql.DB
}

// seederの実行。
func seed(s Seed, seedMethodName string) {
	// Get the reflect value of the method
	seederMethod := reflect.ValueOf(s).MethodByName(seedMethodName)

	// seederMethod内に指定された関数が存在しなければ処理を実行しない旨を表示。
	if !seederMethod.IsValid() {
		log.Fatal("No method called ", seedMethodName)
	}
	// 存在していれば、処理を実行。
	log.Println("Seeding", seedMethodName, "...")
	seederMethod.Call(nil)
	log.Println("Seed", seedMethodName, "succedd")
}

/**
 * 受け取ったseedMethodNamesを実行。
 * もし、seedMethodNamesが指定されていなかった場合は、全てのseedersを実行する。
 */
func Execute(db *sql.DB, seedMethodNames ...string) {
	s := Seed{db}
	seedType := reflect.TypeOf(s)
	log.Println("Running all seeder...", seedType.NumMethod())
	// seedMethodNamesに名前がセットされていない場合は、seedパッケージのものをすべて実行。
	if len(seedMethodNames) == 0 {
		log.Println("Running all seeder...")

		// seedType.NumMethod()でtype Seed structに紐付けられているメソッド数を取得し、その数分処理を実行。
		for i := 0; i < seedType.NumMethod(); i++ {
			// Get the method in the current iteration
			method := seedType.Method(i)
			// Execute seeder
			seed(s, method.Name)
		}
	}

	// seedMethodNamesに値があった場合は、セットされている分実行する。
	for _, item := range seedMethodNames {
		log.Println("Running all seeder...")
		seed(s, item)
	}
}
