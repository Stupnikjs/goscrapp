package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Race struct {
	Name        string    `json:"name"`
	Date        time.Time `json:"date"`
	City        string    `json:"city"`
	Link        string    `json:"link"`
	Departement int       `json:"dep"`
	Site        string    `json:"site"`
}

func (r *Race) IsComplete() bool {
	if r.Name != "" {
		return true

	}
	return false

}

func RacesToJson(dist *os.File, races []Race) {

	JsonByte, err := json.Marshal(races)

	if err != nil {

		fmt.Println(err)
	}

	_, err = dist.Write(JsonByte)

	if err != nil {

		fmt.Println(err)
	}

	defer dist.Close()

}
