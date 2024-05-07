package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func GetMonthMap() map[string]int {

	mois := []string{"Janvier", "Février", "Mars", "Avril", "Mai", "Juin", "Juillet", "Août", "Septembre", "Octobre", "Novembre", "Décembre"}
	monthIndexMap := make(map[string]int)

	// Populate the map with French month names and their respective index values
	for i, moisItem := range mois {
		monthIndexMap[moisItem] = i
	}

	return monthIndexMap

}

func RaceArrayJson(file *os.File, raceArr *[]Race) error {

	jsonByte, err := json.Marshal(raceArr)

	if err != nil {
		fmt.Println(err)
		return err
	}
	_, err = file.Write(jsonByte)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil

}
