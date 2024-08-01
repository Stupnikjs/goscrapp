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

func NewOcpAnnonce(url string) *Annonce {

	var date, jobtype, employementType, location string

	ctx, _ := chromedp.NewContext(context.Background())
	ctx, cancel := context.WithTimeout(ctx, time.Minute*3)
	defer cancel()

	jobtypeSelector := `//article//h2`
	employementTypeSelector := `//li[@class='job_contract_type']/strong`
	locationSelector := `//article//h3`

	err := chromedp.Run(
		ctx,
		chromedp.Navigate(url),
		chromedp.Text(jobtypeSelector, &jobtype, chromedp.NodeVisible),
		chromedp.Text(employementTypeSelector, &employementType, chromedp.NodeVisible),
		chromedp.Text(locationSelector, &location, chromedp.NodeVisible),
	)

	if err != nil {
		fmt.Println(err)
	}

	return &Annonce{
		Url:        url,
		PubDate:    date,
		Lieu:       location,
		Profession: jobtype,
		Contrat:    employementType,
	}

}

func ScrapOcpUrls(selector string, URL string, nodes []*cdp.Node, urls *[]string) {
	ctx, _ := chromedp.NewContext(context.Background())
	ctx, cancel := context.WithTimeout(ctx, time.Minute*3)
	defer cancel()

	err := chromedp.Run(
		ctx,
		chromedp.Navigate(URL),
		chromedp.Nodes(selector, &nodes),
		chromedp.ActionFunc(func(ctx context.Context) error {
			ProcessOcpNodes(nodes, urls)
			return nil
		}),
	)

	if err != nil {
		log.Fatal(err)
	}

}

func ProcessOcpNodes(nodes []*cdp.Node, urls *[]string) {

	for _, node := range nodes {
		if node.NodeType == cdp.NodeTypeElement {
			*urls = append(*urls, node.Attributes[1])
		}
	}

}

func GetOcpUrls() {
	var nodes []*cdp.Node
	var selector string = `//div[@class="offers"]//*//a/@href`
	var url string = "https://www.petitesannonces-ocp.fr/annonces/offres-emploi"
	var urls = []string{}

	// recuperer le nombres de pages en scrappant

	for i := range [10]int{} {

		if i != 0 {
			url = fmt.Sprintf("https://www.petitesannonces-ocp.fr/annonces/offres-emploi?page=%d", i+1)
		}

		ScrapUrls(selector, url, nodes, &urls)

	}

	bytes, err := json.Marshal(urls)

	if err != nil {
		fmt.Println(err)
	}
	file, err := os.Create("ocpurls.json")

	if err != nil {
		fmt.Println(err)
	}

	_, err = file.Write(bytes)

	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

}
