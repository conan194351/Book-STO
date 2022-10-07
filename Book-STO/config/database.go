package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func getDatabase() *sql.DB {
	err := godotenv.Load("././.env")
	if err != nil {
		fmt.Printf("Error loading .env file")
	}
	dbDriver := "mysql"
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp(127.0.0.1:3306)/"+dbName)
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Database loaded successfully")
	}
	return db
}

func InitDatabase() {
	DB = getDatabase()
}
