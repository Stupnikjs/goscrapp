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
