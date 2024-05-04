package main

import (
)

func getDateInfo(node *cdp.Node) (time.Time,error) {
   var day,mounth,year int 
   if node.Parent == nil {
 return "", error.New("no parent in this node")
}
   if node.Parent.NodeName == "EM" {
    yearStr := node.NodeValue
                                year = strconv.Atoi(yearStr)
                        }
                        if node.Parent.NodeName == "SPAN" {
                                dayStr := node.NodeValue
    day = strconv.Atoi(dayStr)
                        }

                        if node.Parent.NodeName == "STRONG" {
                                monthStr := node.NodeValue
    // process 
    mois := []string{"Janvier", "Février", "Mars", "Avril", "Mai", "Juin", "Juillet", "Août", "Septembre", "Octobre", "Novembre", "Décembre"}
   monthIndexMap := make(map[string]int)

    // Populate the map with French month names and their respective index values
    for i, moisItem := range mois {
        monthIndexMap[moisItem] = i
    }
    mounth = monthIndexMap[monthStr]

}
    return time.Date(year,time.Month(month),day,0,0,0,0,Time.UTC)
                        }