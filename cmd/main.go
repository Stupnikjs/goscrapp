package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

func main() {

	var nodes []*cdp.Node
	var selector string = `//ul[@class="tablelike"]//a/@href`
	var url string = "https://www.lemoniteurdespharmacies.fr/emploi/espace-candidats/lire-les-annonces.html"

	// recuperer le nombres de pages en scrappant
	for i := range [17]int{} {

		urls := Scrap(selector, url, nodes, i)
		fmt.Println(urls)

	}

}

func Scrap(selector string, URL string, nodes []*cdp.Node, i int) *[]string {
	ctx, _ := chromedp.NewContext(context.Background())
	ctx, cancel := context.WithTimeout(ctx, time.Minute*3)
	defer cancel()
	var urls = []string{}
	var url string = ""
	if i == 0 {
		url = "https://www.lemoniteurdespharmacies.fr/emploi/espace-candidats/lire-les-annonces.html"
	}
	url = fmt.Sprintf("https://www.lemoniteurdespharmacies.fr/emploi/espace-candidats/lire-les-annonces-%d.html", i)

	err := chromedp.Run(
		ctx,
		chromedp.Navigate(url),
		chromedp.Nodes(selector, &nodes),
		chromedp.ActionFunc(func(ctx context.Context) error {

			ProcessNodes(nodes, &urls)
			return nil
		}),
	)

	if err != nil {
		log.Fatal(err)
	}
	return &urls
}

func ProcessNodes(nodes []*cdp.Node, urls *[]string) *[]string {

	for _, node := range nodes {
		if node.NodeType == cdp.NodeTypeElement {
			*urls = append(*urls, node.Attributes[1])
		}
	}
	return urls
}
