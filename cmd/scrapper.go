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

func GetCityInfo(node *cdp.Node) (map[string]string, error) {
	cityMap := make(map[string]string)

	firstsplit := strings.Split(node.NodeValue, ")")

	secsplit := strings.Split(firstsplit[0], ")")
	thirdsplit := strings.Split(secsplit[0], "(")
	if len(thirdsplit) > 1 {

		cityMap["departement"] = strings.Split(secsplit[0], "(")[1]
	}
	cityMap["city"] = strings.TrimSpace(firstsplit[0])

	return cityMap, nil
}
func GetDateInfo(node *cdp.Node) (map[string]int, error) {

	dateMap := make(map[string]int)

	if node.Parent == nil {
		return nil, errors.New("no parent in this node")
	}
	if node.Parent.NodeName == "EM" {
		yearStr := node.NodeValue
		year, err := strconv.Atoi(yearStr)
		dateMap["year"] = year
		if err != nil {
			return nil, err
		}
	}
	if node.Parent.NodeName == "SPAN" {
		dayStr := node.NodeValue
		day, err := strconv.Atoi(dayStr)
		dateMap["day"] = day
		if err != nil {
			return nil, err
		}
	}

	if node.Parent.NodeName == "STRONG" {
		monthStr := node.NodeValue

		monthIndexMap := GetMonthMap()
		month := monthIndexMap[monthStr]
		dateMap["month"] = month

	}
	return dateMap, nil
}

func RecurseNodes(w io.Writer, nodes []*cdp.Node, races *[]data.Race, race *data.Race) {
	// This will block until the chromedp listener closes the channel

	for _, node := range nodes {
		if node.NodeName == "#text" && node.Parent.AttributeValue("class") == "Cuprum" {
			race.Name = node.NodeValue
		}

		if node.NodeName == "#text" && node.Parent.Parent.AttributeValue("class") == "col-md-12 textleft" && node.Parent.NodeName == "P" {
			mapCity, err := GetCityInfo(node)
			if err != nil {
				fmt.Println(err)
			}
			race.City = mapCity["city"]
			depStr := mapCity["departement"]

			dep, err := strconv.Atoi(depStr)
			if err != nil {
				fmt.Println(err)
			}
			race.Departement = dep

			if err != nil {
				fmt.Println(err)
			}
		}
		if strings.Contains(node.Parent.Parent.AttributeValue("id"), "calendar") {
			d, err := GetDateInfo(node)
			if err != nil {
				fmt.Println(err)
			}
			race.Date = time.Date(d["year"], time.Month(d["month"]), d["day"], 0, 0, 0, 0, time.UTC).String()
		}

		if race.IsFull() && !race.IsInRaces(races) {
			*races = append(*races, *race)
			RecurseNodes(w, node.Children, races, race)
		}

		if node.ChildNodeCount > 0 {
			RecurseNodes(w, node.Children, races, race)

		}

	}

}
