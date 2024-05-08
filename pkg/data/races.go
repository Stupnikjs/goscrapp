package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Race struct {
	Name        string `json:"name"`
	Date        string `json:"date"`
	City        string `json:"city"`
	Departement int    `json:"dep"`
	/*
		Link        string `json:"link"`

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
	if r.Name != "" && r.Date != "" && r.City != "" {
		return true
	}
	return false
}

func (r *Race) isInRaces(arr *[]Race) bool {
	r_val := *r
	races := *arr
	for _, race := range races {
		if r_val.Name == race.Name && r_val.Date == race.Date && r_val.City == race.City && r_val.Departement == race.Departement {
			return true
		}
	}
	return false
}
