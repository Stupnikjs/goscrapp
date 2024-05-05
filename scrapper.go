package main

import (
        "strings"
        "strconv"
        "io"
        "errors"
        "github.com/chromedp/cdproto/cdp"
        "github.com/chromedp/chromedp"
)

func GetDateInfo(node *cdp.Node) (time.Time,error) {
   var day,mounth,year int 
   if node.Parent == nil {
 return "", errors.New("no parent in this node")
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
   
    
    monthIndexMap := GetMonthMap()
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
    
    d,err := GetDateInfo(node)  
    if err := nil {
  fmt.Println(err)
}
    race.Day = d
			}

  // pas de noeuds enfant 
		if node.ChildNodeCount > 0 {
			printNodes(w, node.Children, race)
		}
		
  
	}

}


