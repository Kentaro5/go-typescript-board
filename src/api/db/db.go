package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// NewConnection creates the connection to the database
func NewConnection() (*sql.DB, error) {
	filePath, err := filepath.Abs(".env")
	if err != nil {
		fmt.Println(err)
	}

	err = godotenv.Load(fmt.Sprintf(filePath))
	if err != nil {
		log.Fatalf("godotenvが使用できません。godotenvをロードしてください。", err)
	}

	connectionString := fmt.Sprintf("%v:%v@(%v:%v)/%v",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	// connect DB
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatalf("Error opening DB: %v", err)
		return nil, err
	}

	return db, nil
}
