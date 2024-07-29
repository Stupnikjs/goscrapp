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
	var urls = []string{}
	// recuperer le nombres de pages en scrappant
	for i := range [3]int{} {

		if i != 0 {
			url = fmt.Sprintf("https://www.lemoniteurdespharmacies.fr/emploi/espace-candidats/lire-les-annonces-%d.html", i)
		}

		Scrap(selector, url, nodes, &urls)

	}

}

func Scrap(selector string, URL string, nodes []*cdp.Node, urls *[]string) *[]string {
	ctx, _ := chromedp.NewContext(context.Background())
	ctx, cancel := context.WithTimeout(ctx, time.Minute*3)
	defer cancel()

	err := chromedp.Run(
		ctx,
		chromedp.Navigate(URL),
		chromedp.Nodes(selector, &nodes),
		chromedp.ActionFunc(func(ctx context.Context) error {
			ProcessNodes(nodes, urls)
			return nil
		}),
	)

	if err != nil {
		log.Fatal(err)
	}
	return urls
}

func ProcessNodes(nodes []*cdp.Node, urls *[]string) *[]string {
	newUrls := *urls
	for _, node := range nodes {
		if node.NodeType == cdp.NodeTypeElement {
			newUrls = append(newUrls, node.Attributes[1])
		}
	}
	fmt.Println(newUrls)
	return &newUrls
}
