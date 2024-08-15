package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Stupnikjs/goscrapp/database"
	"github.com/Stupnikjs/goscrapp/scrap"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load("./.env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	db, err := ConnectToDB()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("connection succesfull")
	}

	repo := database.PostgresRepo{
		DB: db,
	}
	app := Application{
		DB:       &repo,
		Scrapper: &scrap.Test,
	}

	err = app.DB.InitTable()
	fmt.Println(err)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(">: ")
	for scanner.Scan() {
		app.CommandParser(strings.TrimSpace(scanner.Text()))
		fmt.Print(">: ")
	}

}

func Exit() {
	os.Exit(1)
}
