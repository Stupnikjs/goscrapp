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
func OpenUrls() []string {
	var urls = []string{}
	file, _ := os.Open("moniteururls.json")
	defer file.Close()
	bytes, _ := io.ReadAll(file)
	_ = json.Unmarshal(bytes, &urls)
	fmt.Println("urls len", len(urls))
	return urls
}

func CreateAnnoncesFile() {
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
}
