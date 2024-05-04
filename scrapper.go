package main

import (
        "strings"
        "strconv"
        "io"
        "github.com/chromedp/cdproto/cdp"
        "github.com/chromedp/chromedp"
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


func printNodes(w io.Writer, nodes []*cdp.Node, race *Race) {
	// This will block until the chromedp listener closes the channel

	for _, node := range nodes {

		if node.NodeName == "#text" {

			if node.Parent.AttributeValue("class") == "Cuprum" {
				race.Name = node.NodeValue
				print("name", race.Name)
			}

		}
		if strings.Contains(node.Parent.Parent.AttributeValue("id"), "calendar") {
			if node.Parent.NodeName == "EM" {
				race.Year = node.NodeValue
			}
			if node.Parent.NodeName == "SPAN" {
				race.Day = node.NodeValue
			}

			if node.Parent.NodeName == "STRONG" {
				race.Month = node.NodeValue
			}

			fmt.Println("here", node.NodeValue)

		}
		if node.ChildNodeCount > 0 {
			printNodes(w, node.Children, race)
		}
		fmt.Println(race)
  
	}

}