package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
  "strconv"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

type Race struct {
	Name  string
	Date time.Time
 Link string
 Lieu string 
 Departement int 
 
 
}

func main() {

	// create context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
	var race Race
	// run task list
	var nodes []*cdp.Node
	err := chromedp.Run(

		ctx, chromedp.Tasks{
			chromedp.Navigate("https://protiming.fr/Runnings/liste"),

			chromedp.Nodes(`//div[@class="col-md-6 clickable visible-lg visible-md"]//*`, &nodes),
			chromedp.ActionFunc(func(ctx context.Context) error {

				printNodes(os.Stdout, nodes, &race)
				return nil
			}),
		})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("the end")
	os.Exit(2)
}




