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

var Ocp ScrapperPharma

var OcpSelectors = Selectors{
	LieuSelector:    `//article//h3`,
	EmploiSelector:  `//article//h2`,
	ContratSelector: `//li[@class='job_contract_type']/strong`,
}



func ScrapOcpUrls() {
	var nodes []*cdp.Node
	var selector string = `//div[contains(@class, 'offer') and contains(@class, 'theme_2')]//a/@href`
	var URL string = "https://www.petitesannonces-ocp.fr/annonces/offres-emploi"
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
		fmt.Println(err)
	}

}


func (m *ScrapperPharma) GetOcpUrls() {
	var nodes []*cdp.Node
	var url string = "https://www.petitesannonces-ocp.fr/annonces/offres-emploi"
	var urls = []string{}

	num := GetOcpPaginatorNum(url)

	for i := range make([]int, num) {
		if i != 0 {
			url = fmt.Sprintf("https://www.petitesannonces-ocp.fr/annonces/offres-emploi?page=%d", i+1)
		}
		ScrapOcpUrls()
	}

	m.Urls = urls 

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

func parseDep(str string) int {
	split := strings.Split(str, ",")

	if len(split) < 1 {
		return 0
	}

	for dep := range data.Departements {
		if strings.Contains(dep, strings.TrimSpace(split[1])) && len(dep) == len(strings.TrimSpace(split[1])) {
			fmt.Println(split[1], dep, data.Departements[dep])
			return data.Departements[dep]

		}
	}
	return 0
}



