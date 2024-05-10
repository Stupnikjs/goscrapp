package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Stupnikjs/goscrapper/pkg/data"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

func main() {
	start := time.Now()
	races := make([]data.Race, 0)
	race := &data.Race{}
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

	err = data.RaceArrayJson(file, &races)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(time.Since(start))
}

func oldGetTasks(nodes []*cdp.Node, races *[]data.Race, race *data.Race) *chromedp.Tasks {

	return &chromedp.Tasks{
		chromedp.Nodes(`//div[@class="col-md-6 clickable visible-lg visible-md"]//*`, &nodes),
		chromedp.ActionFunc(func(ctx context.Context) error {

			ProcessNode(os.Stdout, nodes)
			return nil
		}),
	}
}

func GetActions(url string, nodes []*cdp.Node, races *[]data.Race, race *data.Race) []chromedp.Action {

	actions := []chromedp.Action{
		chromedp.Navigate(url),
	}

	for i := 2; i <= 5; i++ {
		xpath := fmt.Sprintf(`(//ul[@class="paginator pagination pagination-sm pull-right"]/child::*)[%d]`, i)
		actions = append(actions, chromedp.Click(xpath), oldGetTasks(nodes, races, race))
	}
	return actions

}
