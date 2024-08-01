package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(">: ")
	for scanner.Scan() {
		CommandParser(strings.TrimSpace(scanner.Text()))
		fmt.Print(">: ")
	}

}

func Exit() {
	os.Exit(1)
}

// get urls from urls.txt file
func OpenMoniteurUrls() []string {
	var urls = []string{}
	file, _ := os.Open("moniteururls.json")
	defer file.Close()
	bytes, _ := io.ReadAll(file)
	_ = json.Unmarshal(bytes, &urls)
	fmt.Println("urls len", len(urls))
	return urls
}
func OpenOcpUrls() []string {
	var urls = []string{}
	file, _ := os.Open("ocpurls.json")
	defer file.Close()
	bytes, _ := io.ReadAll(file)
	_ = json.Unmarshal(bytes, &urls)
	fmt.Println("urls len", len(urls))
	return urls
}

func CreateMoniteurAnnoncesFile() {
	urls := OpenMoniteurUrls()
	annonces := []Annonce{}
	for _, u := range urls {
		annonce := NewAnnonce(u)
		annonces = append(annonces, *annonce)
	}
	file, _ := os.Create("annonces.json")
	defer file.Close()
	bytes, _ := json.Marshal(annonces)
	file.Write(bytes)
}
