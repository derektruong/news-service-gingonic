package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func DBConnect() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "123456"
	dbName := "NEWS_APP"

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass + "@/" + dbName)

	if err != nil {
		panic(err.Error())
	}

	return db
}
