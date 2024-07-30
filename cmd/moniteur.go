package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

type Annonce struct {
	Url         string `json:"url"`
	Lieu        string `json:"lieu"`
	Departement string `json:"Departement"`
	Profession  string `json:"profession"`
}

func NewAnnonce(url string) {

	ctx, _ := chromedp.NewContext(context.Background())
	ctx, cancel := context.WithTimeout(ctx, time.Minute*3)
	defer cancel()

	selector := `//ul[@class='liste_champs']//li/*`
	nodes := []*cdp.Node{}
	err := chromedp.Run(
		ctx,
		chromedp.Navigate(url),
		chromedp.Nodes(selector, &nodes),
		chromedp.ActionFunc(func(ctx context.Context) error {
			ProcessAnnonceNodes(nodes)
			return nil
		}),
	)

	if err != nil {
		fmt.Println(err)
	}

}

func ProcessAnnonceNodes(nodes []*cdp.Node) {
	for _, node := range nodes {
		fmt.Println("atr", node.Attributes)
		if node.NodeName == "span" {
			fmt.Println("content", node.ContentDocument.NodeName)
		}

		fmt.Println("nodeName", node.NodeName)
		itemprop, _ := node.Attribute("itemprop")
		fmt.Println("itemprop", itemprop)

		if itemprop == "jobLocation" {
			for _, n := range node.Children {
				fmt.Println("children", n.NodeValue)
			}
		}
		if node.NodeName == "#text" {
			fmt.Println("nodeValue", node.NodeValue)
			fmt.Println("node value", node.Value)
		}

		if len(node.Children) > 0 {
			ProcessAnnonceNodes(node.Children)
		}

	}
}

func ScrapUrls(selector string, URL string, nodes []*cdp.Node, urls *[]string) {
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

}

func ProcessNodes(nodes []*cdp.Node, urls *[]string) {

	for _, node := range nodes {
		if node.NodeType == cdp.NodeTypeElement {
			*urls = append(*urls, node.Attributes[1])
		}
	}

}

func GetMoniteurUrls() {
	var nodes []*cdp.Node
	var selector string = `//ul[@class="tablelike"]//a/@href`
	var url string = "https://www.lemoniteurdespharmacies.fr/emploi/espace-candidats/lire-les-annonces.html"
	var urls = []string{}

	// recuperer le nombres de pages en scrappant

	for i := range [10]int{} {

		if i != 0 {
			url = fmt.Sprintf("https://www.lemoniteurdespharmacies.fr/emploi/espace-candidats/lire-les-annonces-%d.html", i)
		}

		ScrapUrls(selector, url, nodes, &urls)

	}

	bytes, err := json.Marshal(urls)

	if err != nil {
		fmt.Println(err)
	}
	file, err := os.Create("urls.txt")

	if err != nil {
		fmt.Println(err)
	}

	_, err = file.Write(bytes)

	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

}
