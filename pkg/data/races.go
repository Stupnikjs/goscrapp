package data

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
	Link        string `json:"link"`
	Site        string `json:"site"`
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

func (r *Race) IsFull() bool {
	if r.Name != "" && r.Date != "" && r.City != "" && r.Link != "" && r.Site == "" &&
		r.Departement != 0 {
		return true
	}
	return false
}

func (r *Race) IsInRaces(arr *[]Race) bool {
	r_val := *r
	races := *arr
	for _, race := range races {
		if race.Equals(r_val) {
			return true
		}
	}
	return false
}

func (r Race) Equals(race Race) bool {
	if r.Name == race.Name && r.Date == race.Date && r.City == race.City &&
		r.Departement == race.Departement && r.Link == race.Link && r.Site == race.Site {
		return true
	}
	return false
}
