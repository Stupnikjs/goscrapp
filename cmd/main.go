package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load("./.env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	db, err := ConnectToDB()

	if err != nil {
		fmt.Println("connection successful to sql")
	}
	app := Application{
		Commands: commandsMap,
		DB:       db,
	}

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
