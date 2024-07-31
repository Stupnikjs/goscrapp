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
	PubDate     string `json:"pubdate"`
	Lieu        string `json:"lieu"`
	Departement string `json:"Departement"`
	Profession  string `json:"profession"`
	Contrat     string `json:"contrat"`
}

func scrapUrlsToJson(url string) {

	ctx, _ := chromedp.NewContext(context.Background())
	ctx, cancel := context.WithTimeout(ctx, time.Minute*3)
	defer cancel()

	var urls = []string{}

	err := chromedp.Run(
		ctx,
		chromedp.Navigate(url),
		chromedp.Evaluate(`document.querySelectorAll()`, &urls),
	)

	if err != nil {
		fmt.Println(err)
	}

}

func NewAnnonce(url string) *Annonce {

	var entreprise, date, jobtype, employementType, location string

	ctx, _ := chromedp.NewContext(context.Background())
	ctx, cancel := context.WithTimeout(ctx, time.Minute*3)
	defer cancel()

	entrepriseSelector := `//*[@itemprop='hiringOrganization']//span[@itemprop="name"]`
	dateSelector := `//span[@itemprop='datePosted']`
	jobtypeSelector := `//span[@itemprop='occupationalCategory']`
	employementTypeSelector := `//span[@itemprop='employmentType']`
	locationSelector := `//span[@itemprop='jobLocation']//span`

	err := chromedp.Run(
		ctx,
		chromedp.Navigate(url),
		chromedp.Text(entrepriseSelector, &entreprise, chromedp.NodeVisible),
		chromedp.Text(dateSelector, &date, chromedp.NodeVisible),
		chromedp.Text(jobtypeSelector, &jobtype, chromedp.NodeVisible),
		chromedp.Text(employementTypeSelector, &employementType, chromedp.NodeVisible),
		chromedp.Text(locationSelector, &location, chromedp.NodeVisible),
	)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("entreprise", entreprise)
	fmt.Println("date", date)
	fmt.Println("location", location)
	fmt.Println("jobtype", jobtype)
	fmt.Println(employementType)

	return &Annonce{
		Url:        url,
		PubDate:    date,
		Lieu:       location,
		Profession: jobtype,
		Contrat:    employementType,
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
	file, err := os.Create("urls.json")

	if err != nil {
		fmt.Println(err)
	}

	_, err = file.Write(bytes)

	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

}
