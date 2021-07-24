package db

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"api/seed"
)

func main() {
	handleArgs()
}

func handleArgs() {
	flag.Parse()
	args := flag.Args()
	log.Println("mysql", flag.Parse())
	log.Println("mysql", flag.Args())

	if len(args) >= 1 {
		switch args[0] {
		case "seed":
			connString := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local&multiStatements=true", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
			// connect DB
			db, err := sql.Open("mysql", connString)
			if err != nil {
				log.Fatalf("Error opening DB: %v", err)
			}
			seeds.Execute(db, args[1:]...)
			os.Exit(0)
		}
	}
}
