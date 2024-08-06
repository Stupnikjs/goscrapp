package utils

import (
	"encoding/json"
	"os"
)

func ArrToJson[T any](arr []T, filename string) error {

	bytes, err := json.Marshal(arr)
	if err != nil {
		return err
	}
	f, _ := os.Create(filename)
	f.Write(bytes)
	defer f.Close()
	return nil
}
