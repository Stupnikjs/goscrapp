package main

import (
	"encoding/json"
	"io"
	"os"
)

type Annonce struct {
	Url         string `json:"url"`
	PubDate     string `json:"pubdate"`
	Lieu        string `json:"lieu"`
	Departement int    `json:"departement"`
	Profession  string `json:"profession"`
	Contrat     string `json:"contrat"`
}

func GetAllAnnnonces() []Annonce {

	var urls = []Annonce{}
	file, _ := os.Open("annonces.json")
	defer file.Close()
	bytes, _ := io.ReadAll(file)
	_ = json.Unmarshal(bytes, &urls)
	return urls

}
