package postgresql

import (
	"database/sql"
	"fmt"
	"log"
)

var DB *sql.DB

func Init(host string, dbName string, user string, password string) {
	var err error

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s", host, user, password, dbName)
	DB, err = sql.Open("postgres", dsn)

	if err != nil {
		log.Fatal(err)
	}

	// Open doesn't open a connection. Validate DSN data:
	err = DB.Ping()

	if err != nil {
		log.Fatal(err)
	} else {
		println("Connection established")
	}
}

func Close() {
	DB.Close()
}
