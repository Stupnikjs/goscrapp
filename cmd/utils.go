package main

func GetMonthMap() map[string]int {

	mois := []string{"Janv.", "Févr.", "Mars", "Avril", "Mai", "Juin", "Juil.", "Août", "Sept.", "Oct.", "Nov.", "Déc."}
	monthIndexMap := make(map[string]int)

	// Populate the map with French month names and their respective index values
	for i, moisItem := range mois {
		monthIndexMap[moisItem] = i + 1
	}

	return monthIndexMap

}
