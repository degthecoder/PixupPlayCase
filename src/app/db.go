package app

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var Db *sql.DB

func ConnectDb() {
	conStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		Settings.DbUser,
		Settings.DbPassword,
		Settings.DbHost,
		Settings.DbPort,
		Settings.DbName,
	)

	db, err := sql.Open("mysql", conStr)
	if err != nil {
		log.Fatalf("Could not open DB connection: %v", err)
	}

	db.SetMaxIdleConns(0)
	Db = db
}

func DisconnectDb() {
	if Db != nil {
		Db.Close()
	}
}
