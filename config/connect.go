package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func DbConn() *sql.DB {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "root"
	dbName := "longphu"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}
