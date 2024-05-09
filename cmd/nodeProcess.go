package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Stupnikjs/goscrapper/pkg/data"
	"github.com/chromedp/cdproto/cdp"
)

func GetCityInfo(node *cdp.Node) (map[string]string, error) {
	cityMap := make(map[string]string)

	firstsplit := strings.Split(node.NodeValue, "(")

	if len(firstsplit) != 2 {

		cityMap["city"] = "online"
		cityMap["departement"] = "online"
		return cityMap, nil
	}

	secsplit := strings.Split(firstsplit[1], ")")
	if len(secsplit) > 1 {

		cityMap["departement"] = strings.TrimSpace(secsplit[0])
	}
	cityMap["city"] = strings.TrimSpace(firstsplit[0])

	return cityMap, nil
}

func GetDateInfo(node *cdp.Node) (map[string]int, error) {

	dateMap := make(map[string]int)

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
			dateMap["month"] = month

		}
	}

	return dateMap, nil
}

func NodeQuery(node *cdp.Node, race *data.Race) (*data.Race, error) {

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

	if !race.IsFull() {
		for _, n := range node.Children {
			NodeQuery(n, race)
		}

	}

	return race, nil
}
