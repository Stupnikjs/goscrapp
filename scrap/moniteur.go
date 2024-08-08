package scrap

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

// creer un checkeur de doublons

var MoniteurSelectors = Selectors{
	EntepriseSelector: `//*[@itemprop='hiringOrganization']//span[@itemprop="name"]`,
	DateSelector:      `//span[@itemprop='datePosted']`,
	EmploiSelector:    `//span[@itemprop='occupationalCategory']`,
	ContratSelector:   `//span[@itemprop='employmentType']`,
	LieuSelector:      `//span[@itemprop='jobLocation']//span`,
}

var MoniteurScrapper = ScrapperSite{
	Site:        "moniteur",
	Selectors:   MoniteurSelectors,
	UrlScrapper: ScrappMoniteurUrls,
}

func ScrappMoniteurUrls(s *ScrapperSite) *ScrapperSite {
	var selector string = `//ul[@class="tablelike"]//a/@href`
	var url string = "https://www.lemoniteurdespharmacies.fr/emploi/espace-candidats/lire-les-annonces.html"
	pageNum := ScrapPageNumMoniteur()

	for i := range make([]int, pageNum, 16) {

		if i != 0 {
			url = fmt.Sprintf("https://www.lemoniteurdespharmacies.fr/emploi/espace-candidats/lire-les-annonces-%d.html", i)
		}
		var nodes []*cdp.Node
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
						s.Urls = append(s.Urls, node.Attributes[1])
					}
				}
				return nil
			}),
		)
		if err != nil {
			fmt.Println(err)
		}
	}
	return s

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
		}))

	if err != nil {
		fmt.Println(err)
	}

	return pageNum
}

func ProcessPaginator(nodes []*cdp.Node) int {
	return len(nodes) - 4
}

func ExtractDepartement(str string) int {

	split := strings.Split(str, "(")
	if len(split) > 1 {
		if len(split[1]) >= 2 {
			depStr := split[1][:2]
			dep, err := strconv.Atoi(depStr)
			if err != nil {
				return 0
			}
			return dep
		}
	}

	return 0
}

/*
func (m *ScrapperPharma) WrapperScrappAnnonces() {
	if len(m.Urls) == 0 {
		m.ScrappMoniteurUrls()
		utils.ArrToJson(m.Urls, "moniteur_urls.json")
		fmt.Println(m.Urls)
	}
	annonces := m.ScrappAnnonces()

	err := utils.ArrToJson(annonces, "moniteur.json")
	fmt.Println(err)

}
*/
