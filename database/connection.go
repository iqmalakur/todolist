package database

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func GetConnection() (*sql.DB, error) {
	host := os.Getenv("dbhost")
	if host == "" {
		host = "localhost"
	}

	user := os.Getenv("dbuser")
	if user == "" {
		user = "postgres"
	}

	pass := os.Getenv("dbpass")
	if pass == "" {
		pass = "postgres"
	}

	dbname := os.Getenv("dbname")
	if dbname == "" {
		dbname = "todocible"
	}

	conn := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", user, pass, host, dbname)

	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(time.Hour)

	return db, nil
}