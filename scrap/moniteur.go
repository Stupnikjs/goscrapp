package scrap

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"time"

	"github.com/Stupnikjs/goscrapp/data"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

// creer un checkeur de doublons

func NewMoniteurAnnonce(url string) *data.Annonce {

	var entreprise, date, jobtype, employementType, location string

	ctx, _ := chromedp.NewContext(context.Background())
	ctx, cancel := context.WithTimeout(ctx, time.Second*7)
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

	return &data.Annonce{
		Url:        url,
		PubDate:    date,
		Lieu:       location,
		Profession: jobtype,
		Contrat:    employementType,
		Created_at: time.Now().GoString(),
	}

}

func ScrapUrls(selector string, URL string, nodes []*cdp.Node, urls *[]string) {
	ctx, _ := chromedp.NewContext(context.Background())
	ctx, cancel := context.WithTimeout(ctx, time.Second*20)
	defer cancel()

	err := chromedp.Run(
		ctx,
		chromedp.Navigate(URL),
		chromedp.Nodes(selector, &nodes),
		chromedp.ActionFunc(func(ctx context.Context) error {
			for _, node := range nodes {
				if node.NodeType == cdp.NodeTypeElement {
					*urls = append(*urls, node.Attributes[1])
				}
			}
			return nil
		}),
	)

	if err != nil {
		log.Fatal(err)
	}

}

func GetMoniteurUrls() {
	var nodes []*cdp.Node
	var selector string = `//ul[@class="tablelike"]//a/@href`
	var url string = "https://www.lemoniteurdespharmacies.fr/emploi/espace-candidats/lire-les-annonces.html"
	var urls = []string{}

	// recuperer le nombres de pages en scrappant
	pageNum := ScrapPageNumMoniteur()
	for i := range make([]int, pageNum, 16) {
		if i != 0 {
			url = fmt.Sprintf("https://www.lemoniteurdespharmacies.fr/emploi/espace-candidats/lire-les-annonces-%d.html", i)
		}

		ScrapUrls(selector, url, nodes, &urls)

	}

	bytes, err := json.Marshal(urls)

	if err != nil {
		fmt.Println(err)
	}
	file, err := os.Create("moniteururls.json")

	if err != nil {
		fmt.Println(err)
	}

	_, err = file.Write(bytes)

	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

}

func ScrapPageNumMoniteur() int {

	URL := "https://www.lemoniteurdespharmacies.fr/emploi/espace-candidats/lire-les-annonces.html"
	ctx, _ := chromedp.NewContext(context.Background())
	ctx, cancel := context.WithTimeout(ctx, time.Second*20)
	nodes := []*cdp.Node{}
	defer cancel()
	var pageNum int
	selector := `//ul[@id="liste_pagination"]//*`
	err := chromedp.Run(
		ctx,
		chromedp.Navigate(URL),
		chromedp.Nodes(selector, &nodes),
		chromedp.ActionFunc(func(ctx context.Context) error {
			pageNum = ProcessPaginator(nodes)
			return nil
		}),
	)
	if err != nil {
		fmt.Println(err)
	}

	return pageNum
}

func ProcessPaginator(nodes []*cdp.Node) int {
	return len(nodes) - 4
}
