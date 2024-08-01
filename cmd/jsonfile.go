package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/Stupnikjs/goscrapp/data"
	"github.com/Stupnikjs/goscrapp/scrap"
)

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
	annonces := []data.Annonce{}
	for i, u := range urls {
		annonce := scrap.NewMoniteurAnnonce(u)
		fmt.Println(i)
		annonces = append(annonces, *annonce)
	}
	file, _ := os.Create("annonces.json")
	defer file.Close()
	bytes, _ := json.Marshal(annonces)
	file.Write(bytes)
}

func CreateOcpAnnoncesFile() {

	urls := OpenOcpUrls()
	annonces := []data.Annonce{}
	for _, u := range urls {
		annonce := scrap.NewOcpAnnonce(u)
		fmt.Println(annonce)
		annonces = append(annonces, *annonce)
	}
	file, _ := os.Create("ocp_annonces.json")
	defer file.Close()
	bytes, _ := json.Marshal(annonces)
	file.Write(bytes)

}

func MeltJsonAnnonces() {
	finalname := "annonces.json"
	filename := "moniteur_annonces.json"
	file2name := "ocp_annonces.json"

	annonces := []data.Annonce{}
	annonces2 := []data.Annonce{}

	file1, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer os.Remove(filename)

	file2, err := os.Open(file2name)
	if err != nil {
		fmt.Println(err)
	}
	defer os.Remove(file2name)
	byte_f1, err := io.ReadAll(file1)
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(byte_f1, &annonces)
	if err != nil {
		fmt.Println(err)
	}

	byte_f2, err := io.ReadAll(file2)
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(byte_f2, &annonces2)
	if err != nil {
		fmt.Println(err)
	}
	annonce_final := append(annonces, annonces2...)
	bytes, err := json.Marshal(annonce_final)
	if err != nil {
		fmt.Println(err)
	}
	final, err := os.Create(finalname)

	if err != nil {
		fmt.Println(err)
	}

	defer final.Close()
	final.Write(bytes)

}
