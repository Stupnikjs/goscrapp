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
	start := time.Now()
	var nodes []*cdp.Node
	var selector string = `//div[@class="col-md-6 clickable visible-lg visible-md"]//*`
	var url string = "https://protiming.fr/Runnings/liste"

	ctx, _ := chromedp.NewContext(context.Background())
	ctx, cancel := context.WithTimeout(ctx, time.Minute*3)
	defer cancel()

	Scrap(ctx, selector, url, nodes)

	fmt.Println(time.Since(start))
}

func Scrap(ctx context.Context, selector string, URL string, nodes []*cdp.Node) {
	obj := map[string]string{}
	err := chromedp.Run(
		ctx,
		chromedp.Navigate(URL),
		&chromedp.Tasks{
			chromedp.Nodes(selector, &nodes),
			chromedp.ActionFunc(func(ctx context.Context) error {

				ProcessNode(os.Stdout, nodes, &obj)
				return nil
			}),
		},
	)

	if err != nil {
		log.Fatal(err)
	}

}
