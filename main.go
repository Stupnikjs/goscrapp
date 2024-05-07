package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

func main() {
	races := make([]Race, 0)
	race := &Race{}
	var nodes []*cdp.Node
	// create context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	url := "https://protiming.fr/Runnings/liste"
	err := chromedp.Run(
		ctx,
		chromedp.Navigate(url),
		getTasks(nodes, &races, race),
		chromedp.Click(`(//ul[@class="paginator pagination pagination-sm pull-right"]/child::*)[2]`),
		getTasks(nodes, &races, race),
		chromedp.Click(`(//ul[@class="paginator pagination pagination-sm pull-right"]/child::*)[3]`),
	)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(races)
}

func getTasks(nodes []*cdp.Node, races *[]Race, race *Race) *chromedp.Tasks {

	return &chromedp.Tasks{

		chromedp.Nodes(`//div[@class="col-md-6 clickable visible-lg visible-md"]//*`, &nodes),
		chromedp.ActionFunc(func(ctx context.Context) error {

			printNodes(os.Stdout, nodes, races, race)
			return nil
		}),
	}

}
