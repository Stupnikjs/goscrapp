package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func CommandParser(cmd string) {
	switch cmd {
	case "exit":
		os.Exit(1)
	case "s":
		urls := OpenUrls()
		annonces := []Annonce{}
		for _, u := range urls {

			annonce := NewAnnonce(u)
			annonces = append(annonces, *annonce)

		}
		file, _ := os.Create("annonces.json")
		defer file.Close()
		bytes, _ := json.Marshal(annonces)
		file.Write(bytes)

	case "murl":
		GetMoniteurUrls()
	case "ocpurl":
		GetOcpUrls()
	case "serve":
		Server()
	default:
		fmt.Println("unknown command")

	}
}
