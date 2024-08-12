package scrap

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/cloudsqlconn"
	"cloud.google.com/go/cloudsqlconn/postgres/pgxv4"
)

func openDB() (*sql.DB, error) {

	cleanup, err := pgxv4.RegisterDriver("cloudsql-postgres", cloudsqlconn.WithCredentialsFile("credentials.json"))
	if err != nil {
		fmt.Println(err)
	}
	// call cleanup when you're done with the database connection
	defer cleanup()

	if err != nil {
		fmt.Println(err)

	}

	// Call cleanup when you're done with the database connection

	db, err := sql.Open(
		"cloudsql-postgres",
		fmt.Sprintf("host=%s user=postgres password=%s dbname=postgres sslmode=disable", os.Getenv("SQL_HOST"), os.Getenv("SQL_PASSWORD")))

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		fmt.Print(err)
		return nil, err
	}

	return db, nil
}

func ConnectToDB() (*sql.DB, error) {

	connection, err := openDB()

	if err != nil {
		return nil, err
	}
	log.Println("Connected to Postgres!")
	return connection, nil
}
