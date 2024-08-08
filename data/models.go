package data

import (
	"encoding/json"
	"io"
	"os"
)

type Annonce struct {
	Id          string `json:"id"`
	Digest      string `json:"digest"`
	Url         string `json:"url"`
	PubDate     string `json:"pubdate"`
	Ville       string `json:"ville"`
	Lieu        string `json:"lieu"`
	Departement int    `json:"departement"`
	Description string `json:"description"`
	Profession  string `json:"profession"`
	Contrat     string `json:"contrat"`
	Created_at  string `json:"created_at"`
}

func GetAllAnnnonces() []Annonce {

	var annonces = []Annonce{}
	file, _ := os.Open("annonces.json")
	defer file.Close()
	bytes, _ := io.ReadAll(file)
	_ = json.Unmarshal(bytes, &annonces)
	return annonces

}

func GetAnnoncesHash() {

}
