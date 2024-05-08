package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

func main() {
	races := make([]Race, 0)
	race := &Race{}
	var nodes []*cdp.Node

	ctx, _ := chromedp.NewContext(context.Background())
	ctx, cancel := context.WithTimeout(ctx, time.Minute*3)
	defer cancel()

	url := "https://protiming.fr/Runnings/liste"

	actions := GetActions(url, nodes, &races, race)

	err := chromedp.Run(
		ctx,
		actions...,
	)

	if err != nil {
		log.Fatal(err)
	}

	file, err := os.OpenFile("new.json", os.O_CREATE, 0664)

	if err != nil {
		log.Fatal(err)
	}

	err = RaceArrayJson(file, &races)

	if err != nil {
		log.Fatal(err)
	}
}

func oldGetTasks(nodes []*cdp.Node, races *[]Race, race *Race) *chromedp.Tasks {

	return &chromedp.Tasks{
		chromedp.Nodes(`//div[@class="col-md-6 clickable visible-lg visible-md"]//*`, &nodes),
		chromedp.ActionFunc(func(ctx context.Context) error {

			RecurseNodes(os.Stdout, nodes, races, race)
			return nil
		}),
	}
}

func GetActions(url string, nodes []*cdp.Node, races *[]Race, race *Race) []chromedp.Action {

	actions := []chromedp.Action{
		chromedp.Navigate(url),
	}

	for i := 2; i <= 5; i++ {
		xpath := fmt.Sprintf(`(//ul[@class="paginator pagination pagination-sm pull-right"]/child::*)[%d]`, i)
		actions = append(actions, chromedp.Click(xpath), oldGetTasks(nodes, races, race))
	}
	return actions

}
