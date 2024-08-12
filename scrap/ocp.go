package scrap

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

var OcpSelectors = []Selector{
	{SelectorPath: `//article//h3`, Name: "lieu"},
	{SelectorPath: `//article//h2`, Name: "emploi"},
	{SelectorPath: `//li[@class='job_contract_type']/strong`, Name: "contrat"},
}

var OcpScrapper = ScrapperSite{
	Site:        "ocp",
	Selectors:   OcpSelectors,
	UrlScrapper: GetOcpUrls,
}

func ScrapOcpUrls(url string) []string {
	var nodes []*cdp.Node
	var urls []string
	var selector string = `//div[contains(@class, 'offer') and contains(@class, 'theme_2')]//a/@href`
	ctx, _ := chromedp.NewContext(context.Background())
	ctx, cancel := context.WithTimeout(ctx, time.Second*20)
	defer cancel()

	err := chromedp.Run(
		ctx,
		chromedp.Navigate(url),
		chromedp.Nodes(selector, &nodes),
		chromedp.ActionFunc(func(ctx context.Context) error {
			for _, node := range nodes {
				if node.NodeType == cdp.NodeTypeElement {
					urls = append(urls, node.Attributes[1])
				}
			}
			return nil
		}),
	)

	if err != nil {
		fmt.Println(err)
	}

	return urls
}

func GetOcpUrls(s *ScrapperSite) *ScrapperSite {
	if s.Site != "ocp" {
		fmt.Println("wrong site called")
		return s
	}
	var url string = "https://www.petitesannonces-ocp.fr/annonces/offres-emploi"

	num := GetOcpPaginatorNum(url)

	for i := range make([]int, num) {
		if i != 0 {
			url = fmt.Sprintf("https://www.petitesannonces-ocp.fr/annonces/offres-emploi?page=%d", i+1)
		}
		s.Urls = append(s.Urls, ScrapOcpUrls(url)...)
	}
	return s
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
