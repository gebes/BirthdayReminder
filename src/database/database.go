package database

import (
	"database/sql"
	"gebes.io/go-birthdayreminder/src/env"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

var (
	Database *sql.DB
)

func InitDatabase() {
	log.Println("Connecting to database")
	db, err := sql.Open("mysql", env.MySqlDatabase)
	if err != nil {
		log.Fatalln(err)
	}
	Database = db
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)

	// ping the database, just to be sure
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Successfully connected to the database")
	log.Printf("Currently %d connections are open\n", db.Stats().OpenConnections)
}

func close(countResult *sql.Rows) {
	if countResult == nil {
		return
	}
	err := countResult.Close()
	if err != nil {
		log.Println("Error closing", err)
	}
}
