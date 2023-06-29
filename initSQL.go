package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

// Initializes SQL database
func initSQL() {
	cfg := mysql.Config{
		User:   "root",
		Passwd: "pass",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "tasks",
	}

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
}
