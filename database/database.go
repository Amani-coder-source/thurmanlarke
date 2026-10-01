package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() {

	dsn := "root:Rich$com8a@tcp(127.0.0.1:3306)/thurmanlark"

	var err error

	DB, err = sql.Open("mysql", dsn)

	if err != nil {
		log.Fatal("Error opening database:", err)
	}

	err = DB.Ping()

	if err != nil {
		log.Fatal("Error connecting to MySQL:", err)
	}

	fmt.Println("Successfully connected to Thurmanlark database")
}
