package data

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// get urls from urls.txt file
func OpenUrls(filename string) []string {
	var urls = []string{}
	file, _ := os.Open(filename)
	defer file.Close()
	bytes, _ := io.ReadAll(file)
	_ = json.Unmarshal(bytes, &urls)
	fmt.Println("urls len", len(urls))
	return urls[:20]
}

func CreateJsonFileFromAnnonce(filename string, annonces []Annonce) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	bytes, err := json.Marshal(annonces)
	if err != nil {
		return err
	}

	file.Write(bytes)

	return nil

}

// get called by the command parser
func MeltJsonAnnonces() {
	finalname := "annonces.json"
	filename := "moniteur_annonces.json"
	file2name := "ocp_annonces.json"

	annonces := []Annonce{}
	annonces2 := []Annonce{}

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
}
