package scrap

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Stupnikjs/goscrapp/data"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

var OcpSelectors = Selectors{

	LieuSelector:    `//article//h3`,
	EmploiSelector:  `//article//h2`,
	ContratSelector: `//li[@class='job_contract_type']/strong`,
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

func (s *ScrapperSite) ParseDep(str string) int {
	if s.Site == "moniteur" {
		return ExtractDepartement(str)
	}
	split := strings.Split(str, ",")

	if len(split) < 2 {
		return 0
	}

	for dep := range data.Departements {
		if strings.Contains(dep, strings.TrimSpace(split[1])) && len(dep) == len(strings.TrimSpace(split[1])) {
			return data.Departements[dep]

		}
	}
	return 0
}

func TestOcpScrapper() {
	OcpScrapper.UrlScrapper(&OcpScrapper)
	fmt.Println(OcpScrapper.Urls)
	for _, url := range OcpScrapper.Urls[:10] {
		OcpScrapper.GetAnnonce(url)

	}
	fmt.Println(OcpScrapper.Annonces)

}
