package lib

import (
	"database/sql"
	"log"
	"os"
)

var DB *sql.DB

func ConnectToDb() *sql.DB {
	// connect to the databse
	//db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	log.Print(os.Getenv("ENV"))
	db, err := sql.Open("postgres", "postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable")

	if err != nil {
		log.Fatal("1", err)
	}

	//create table if not exists
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        firstname TEXT,
        lastname TEXT,
        username TEXT UNIQUE,
        password TEXT,
        email TEXT UNIQUE
    );
`)

	if err != nil {
		log.Fatal("2", err)
	}

	DB = db
	return db
}
