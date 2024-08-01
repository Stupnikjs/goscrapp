package scrap

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/Stupnikjs/goscrapp/data"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

func NewOcpAnnonce(url string) *data.Annonce {

	var date, jobtype, employementType, location, description string

	ctx, _ := chromedp.NewContext(context.Background())
	ctx, cancel := context.WithTimeout(ctx, time.Second*8)
	defer cancel()

	jobtypeSelector := `//article//h2`
	employementTypeSelector := `//li[@class='job_contract_type']/strong`
	locationSelector := `//article//h3`
	descriptionSelector := `//p[@id="description"]`

	err := chromedp.Run(
		ctx,
		chromedp.Navigate(url),
		chromedp.Text(jobtypeSelector, &jobtype, chromedp.NodeVisible),
		chromedp.Text(employementTypeSelector, &employementType, chromedp.NodeVisible),
		chromedp.Text(locationSelector, &location, chromedp.NodeVisible),
		chromedp.Text(descriptionSelector, &description, chromedp.NodeVisible),
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
	}

}

func ScrapOcpUrls(selector string, URL string, nodes []*cdp.Node, urls *[]string) {
	ctx, _ := chromedp.NewContext(context.Background())
	ctx, cancel := context.WithTimeout(ctx, time.Second*20)
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
	var selector string = `//div[contains(@class, 'offer') and contains(@class, 'theme_2')]//a/@href`
	var url string = "https://www.petitesannonces-ocp.fr/annonces/offres-emploi"
	var urls = []string{}

	// recuperer le nombres de pages en scrappant

	num := GetOcpPaginatorNum(url)

	for i := range make([]int, num) {
		if i != 0 {
			url = fmt.Sprintf("https://www.petitesannonces-ocp.fr/annonces/offres-emploi?page=%d", i+1)
		}
		print("here")
		ScrapOcpUrls(selector, url, nodes, &urls)
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

func GetOcpPaginatorNum(url string) int {

	ctx, _ := chromedp.NewContext(context.Background())
	ctx, cancel := context.WithTimeout(ctx, time.Second*15)
	defer cancel()

	selectorPaginator := `//li[@class="last_page"]//a/@href`
	var pageNum string
	var pagenodes []*cdp.Node
	err := chromedp.Run(
		ctx,
		chromedp.Navigate(url),
		chromedp.Nodes(selectorPaginator, &pagenodes),
		chromedp.ActionFunc(func(ctx context.Context) error {
			for _, n := range pagenodes {
				if len(n.Attributes) > 0 {
					pageNum = n.Attributes[1]
					fmt.Println(n)
				}
			}
			return nil
		}),
	)

	if err != nil {
		fmt.Println(err)
	}
	strPage := pageNum[len(pageNum)-2:]
	pageInt, err := strconv.Atoi(strPage)
	if err != nil {
		fmt.Println(err)
	}

	return pageInt
}
