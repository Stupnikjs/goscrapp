package data

import (
	"encoding/json"
	"fmt"
	"os"
)

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

func RemoveDuplicates(raceArr *[]Race) []Race {
	arr := *raceArr
	returnArr := []Race{}
	for _, race := range arr {
		inreturnarr := false
		for _, r := range returnArr {
			if race.Equals(r) {
				inreturnarr = true
			}
		}
		if !inreturnarr {
			returnArr = append(returnArr, race)
		}
	}
	return returnArr
}
