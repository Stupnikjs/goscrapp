package main

import (
	"fmt"
	"io"

	"github.com/Stupnikjs/goscrapper/pkg/data"
	"github.com/chromedp/cdproto/cdp"
)

// NOT WORKING

func ProcessNode(w io.Writer, nodes []*cdp.Node, races *[]data.Race, race *data.Race) {

	for i, node := range nodes {
		fmt.Println("node number :", i)
		race, err := NodeQuery(node, race)
		if err != nil {
			fmt.Println(err)
		}

		if race.IsFull() && !race.IsInRaces(races) {
			race.Site = "protiming"
			*races = append(*races, *race)
			race = &data.Race{}

		}

	}

}
