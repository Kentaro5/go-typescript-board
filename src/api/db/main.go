package db

import (
	"api/db/seed"
	"database/sql"
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	handleArgs()
}

func handleArgs() {
	// コマンドラインで指定された、値を使用するために呼ぶ。
	flag.Parse()
	// Parseされた値を配列形式で取得。
	args := flag.Args()
	err := godotenv.Load(fmt.Sprintf("../%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		log.Fatalf("godotenvが使用できません。godotenvをロードしてください。", err)
	}

	if len(args) >= 1 {
		switch args[0] {
		case "seed":
			connString := fmt.Sprintf("%v:%v@(%v:%v)/%v", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))

			log.Println(connString)
			// connect DB
			db, err := sql.Open("mysql", connString)
			if err != nil {
				log.Fatalf("Error opening DB: %v", err)
			}
			seed.Execute(db, args[1:]...)
			os.Exit(0)
		}
	}
}
