package main

import (
	"context"
	"log"
	"os"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

func main() {

	// create context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
	url := "https://protiming.fr/Runnings/liste"
	err := chromedp.Run(
		ctx,
		chromedp.Navigate(url),
		getTasks(),
		chromedp.Click(`(//ul[@class="paginator pagination pagination-sm pull-right"]/child::*)[2]`),
		getTasks(),
		chromedp.Click(`(//ul[@class="paginator pagination pagination-sm pull-right"]/child::*)[3]`),
	)

	if err != nil {
		log.Fatal(err)
	}

	os.Exit(2)
}

func getTasks() *chromedp.Tasks {
	var race Race
	var nodes []*cdp.Node
	return &chromedp.Tasks{

		chromedp.Nodes(`//div[@class="col-md-6 clickable visible-lg visible-md"]//*`, &nodes),
		chromedp.ActionFunc(func(ctx context.Context) error {

			printNodes(os.Stdout, nodes, &race)
			return nil
		}),
	}
}
