package main

import (
	"encoding/json"
	"fmt"
	"os"
)

var commandsMap = map[string]func(){
	"exit": Exit,
	"an":   CreateAnnoncesFile,
	"murl": GetMoniteurUrls,
	"ourl": GetOcpUrls,
	"dep":  ParseLieu,
}

func CommandParser(cmd string) {
	f, ok := commandsMap[cmd]
	if ok {
		f()
	} else {
		fmt.Print("unknown command")
	}
}

func ParseLieu() {
	annonces := GetAllAnnnonces()
	newAnnonces := []Annonce{}
	for _, a := range annonces {
		new := ExtractDepartement(a)
		newAnnonces = append(newAnnonces, new)
	}

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
