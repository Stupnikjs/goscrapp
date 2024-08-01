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
	Url         string `json:"url"`
	PubDate     string `json:"pubdate"`
	Lieu        string `json:"lieu"`
	Region      string `json:"region"`
	Departement int    `json:"departement"`
	Description string `json:"description"`
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
			if strings.Contains(a.Region, dep) {
				a.Departement = Departements[dep]
				fmt.Println(dep)
			}
		}
		newAnnonces = append(newAnnonces, a)

	}
	RemoveOldAnnoncesJson(newAnnonces)

}

func RemoveOldAnnoncesJson(newAnnonces []Annonce) {

	os.Remove("annonces.json")

	newFile, err := os.Create("annonces.json")
	if err != nil {
		fmt.Println(err)
	}
	defer newFile.Close()
	if err != nil {
		fmt.Println(err)
	}
	bytes, err := json.Marshal(newAnnonces)
	if err != nil {
		fmt.Println(err)
	}
	newFile.Write(bytes)
	if err != nil {
		fmt.Println(err)
	}
	newFile.Write(bytes)
}

func ExtractDepartement(a Annonce) Annonce {

	split := strings.Split(a.Region, "(")
	if len(split) > 1 {
		if len(split[1]) >= 2 {
			depStr := split[1][:2]
			dep, err := strconv.Atoi(depStr)
			if err != nil {
				return Annonce{}
			}
			a.Departement = dep
		}
	}

	return a
}
