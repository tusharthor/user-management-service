package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// var DB *sql.DB

func DatabaseConnection() (*sql.DB, error) {
	//by default it loads enf file (load()), also can pass another filename
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("error in loading .env file")
	}

	//entire url with formatted string
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	db, err := sql.Open("mysql", dsn) //dsn - datasourcename
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	log.Println("connected to database")
	return db, err
}
