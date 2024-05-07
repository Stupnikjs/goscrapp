package main

import (
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/chromedp/cdproto/cdp"
)

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

func printNodes(w io.Writer, nodes []*cdp.Node, races *[]Race, race *Race) {
	// This will block until the chromedp listener closes the channel

	for _, node := range nodes {
		if node.NodeName == "#text" && node.Parent.AttributeValue("class") == "Cuprum" {
			race.Name = node.NodeValue
		}

		if strings.Contains(node.Parent.Parent.AttributeValue("id"), "calendar") {
			d, err := GetDateInfo(node)
			if err != nil {
				fmt.Println(err)
			}
			race.Date = time.Date(d["year"], time.Month(d["month"]), d["day"], 0, 0, 0, 0, time.UTC).String()
		}

		if race.isFull() &&  {
			*races = append(*races, *race)
			printNodes(w, node.Children, races, race)
		}

		if node.ChildNodeCount > 0 {
			printNodes(w, node.Children, races, race)

		}

	}

}
