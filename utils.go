package main

func GetMonthMap() map[string]int {

	mois := []string{"Janvier", "Février", "Mars", "Avril", "Mai", "Juin", "Juillet", "Août", "Septembre", "Octobre", "Novembre", "Décembre"}
	monthIndexMap := make(map[string]int)

	// Populate the map with French month names and their respective index values
	for i, moisItem := range mois {
		monthIndexMap[moisItem] = i
	}

	return monthIndexMap

}

func RaceIsInArray(AllRaces *[]*Race, race Race) bool {
	if AllRaces == nil {
		return false
	}

	// Dereference the pointer to access the slice
	races := *AllRaces
	for _, r := range races {
		if r.Date == race.Date && r.Name == race.Name {
			return true
		}
	}
	return false
}
