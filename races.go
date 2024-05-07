package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Race struct {
	Name string `json:"name"`
	Date string `json:"date"`
	/*
		City        string `json:"city"`
		Link        string `json:"link"`
		Departement int    `json:"dep"`
		Site        string `json:"site"`*
	*/
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

func (r *Race) isFull() bool {
	if r.Name != "" && r.Date != "" {
		return true
	}
	return false
}

func (r *Race) isInRaces(arr *[]Race) bool {
	r_val := *r
	races := *arr
	for _, race := range races {
		if r_val.Name == race.Name && r_val.Date == race.Date {
			return true
		}
	}
	return false
}
