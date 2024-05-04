package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
  "strconv"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

type Race struct {
	Name  string
	Day   string
	Month string
	Year  string
}

func main() {

	// create context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
	var race Race
	// run task list
	var nodes []*cdp.Node
	err := chromedp.Run(

		ctx, chromedp.Tasks{
			chromedp.Navigate("https://protiming.fr/Runnings/liste"),

			chromedp.Nodes(`//div[@class="col-md-6 clickable visible-lg visible-md"]//*`, &nodes),
			chromedp.ActionFunc(func(ctx context.Context) error {

				printNodes(os.Stdout, nodes, &race)
				return nil
			}),
		})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("the end")
	os.Exit(2)
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




}