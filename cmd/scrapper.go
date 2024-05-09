package main

import (
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/Stupnikjs/goscrapper/pkg/data"
	"github.com/chromedp/cdproto/cdp"
)

// NOT WORKING
func GetCityInfo(node *cdp.Node) (map[string]string, error) {
	cityMap := make(map[string]string)

	firstsplit := strings.Split(node.NodeValue, "(")

	secsplit := strings.Split(firstsplit[0], ")")

	if len(secsplit) > 1 {

		cityMap["departement"] = strings.TrimSpace(secsplit[0])
	}
	cityMap["city"] = strings.TrimSpace(firstsplit[0])

	return cityMap, nil
}
func GetDateInfo(node *cdp.Node) (map[string]int, error) {

	dateMap := make(map[string]int)
 if node.NodeName != "expected" {
 fmt.Println(node.NodeName)
 return nil, errors.New("wrong node processed by getDateInfo") }

	for _, n := range node.Children {
		
		if n.NodeName == "EM" {
			yearStr := n.Children[0].NodeValue
			year, err := strconv.Atoi(yearStr)

			dateMap["year"] = year
			if err != nil {
				return nil, err
			}
		}
		if n.NodeName == "SPAN" {
			dayStr := n.Children[0].NodeValue
			day, err := strconv.Atoi(dayStr)
			dateMap["day"] = day
			if err != nil {
				return nil, err
			}
		}

		if n.NodeName == "STRONG" {
			monthStr := n.Children[0].NodeValue

			monthIndexMap := GetMonthMap()
			month := monthIndexMap[strings.TrimSpace(monthStr)]
			print(monthStr)
			dateMap["month"] = month

		}
	}
	return dateMap, nil
}

func RecurseNodes(w io.Writer, nodes []*cdp.Node, races *[]data.Race, race *data.Race) {

	// This will block until the chromedp listener closes the channel

	for _, node := range nodes {
		race, err := ProcessNode(node, race)

		if err != nil {
			fmt.Println(err)
		}
		if race.IsFull() && !race.IsInRaces(races) {
			race.Site = "protiming"
			*races = append(*races, *race)
			race = &data.Race{}

			// RecurseNodes(w, node.Children, races, race)
		}

		if node.ChildNodeCount > 0 {
			RecurseNodes(w, node.Children, races, race)

		}

	}

}

func ProcessNode(node *cdp.Node, race *data.Race) (*data.Race, error) {
	if node.NodeName == "#text" && node.Parent.AttributeValue("class") == "Cuprum" {
		race.Name = node.NodeValue
	}

	if node.NodeName == "#text" && node.Parent.Parent.AttributeValue("class") == "col-md-12 textleft" && node.Parent.NodeName == "P" {
		mapCity, err := GetCityInfo(node)
		if err != nil {
			return nil, err
		}
		race.City = mapCity["city"]
		depStr := mapCity["departement"]

		dep, err := strconv.Atoi(depStr)
		if err != nil {
			return nil, err
		}
		race.Departement = dep

	}
	if node.NodeName == "TIME" {
		d, err := GetDateInfo(node)
		if err != nil {
			return nil, err
		}
		timeDate := time.Date(d["year"], time.Month(d["month"]), d["day"], 0, 0, 0, 0, time.UTC)
		race.Date = FormatDate(timeDate)
	}

	if node.AttributeValue("class") == "panel-container event-mosaic-bg" {
		runid := node.AttributeValue("id")[3:]
		race.Link = fmt.Sprintf("https://www.protiming.fr/Runnings/detail/%s", runid)
	}
	return race, nil
}
