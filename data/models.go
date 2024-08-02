package data

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
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

func ParseLieu() {
	annonces := GetAllAnnnonces()
	newAnnonces := []Annonce{}
	for _, a := range annonces {
		new := ExtractDepartement(a)
		newAnnonces = append(newAnnonces, new)
	}
	RemoveOldAnnoncesJson(newAnnonces)
}

func ParseDep() {
	depKey := GetKeys(Departements)
	annonces := GetAllAnnnonces()
	newAnnonces := []Annonce{}
	for _, a := range annonces {
		for _, dep := range depKey {
			if strings.Contains(a.Lieu, dep) {
				a.Departement = Departements[dep]
				fmt.Println(dep)
			}
		}
		newAnnonces = append(newAnnonces, a)

	}
	RemoveOldAnnoncesJson(newAnnonces)

}

func ExtractDepartement(a Annonce) Annonce {

	split := strings.Split(a.Lieu, "(")
	if len(split) > 1 {
		if len(split[1]) >= 2 {
			depStr := split[1][:2]
			dep, err := strconv.Atoi(depStr)
			if err != nil {
				return a
			}
			a.Departement = dep
		}
	}

	return a
}
