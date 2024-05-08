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


func FormatDate(t time.Time) string {
var day,month,year string
// implement this func
day = t.Format("2")   
month = t.Format("01")
year = t.Format("2006")
return fmt.SPrintf("%s%s%s", day,month,year)

}