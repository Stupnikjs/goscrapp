package scrap

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Stupnikjs/goscrapp/data"
	"github.com/Stupnikjs/goscrapp/utils"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

// creer un checkeur de doublons

var MoniteurSelectors = Selectors{
	Site:              "moniteur",
	EntepriseSelector: `//*[@itemprop='hiringOrganization']//span[@itemprop="name"]`,
	DateSelector:      `//span[@itemprop='datePosted']`,
	EmploiSelector:    `//span[@itemprop='occupationalCategory']`,
	ContratSelector:   `//span[@itemprop='employmentType']`,
	LieuSelector:      `//span[@itemprop='jobLocation']//span`,
}

var Moniteur = ScrapperPharma{
	Selectors: MoniteurSelectors,
}

func (m *ScrapperPharma) ScrappAnnonces(sels Selectors) []data.Annonce {
	annonces := []data.Annonce{}

	for _, url := range m.Urls {

		a := m.GetAnnonce(url, sels)
		if a.Url != "" {
			annonces = append(annonces, a)
		}
	}

	return annonces

}

func (m *ScrapperPharma) ScrappMoniteurUrls() {
	var selector string = `//ul[@class="tablelike"]//a/@href`
	var url string = "https://www.lemoniteurdespharmacies.fr/emploi/espace-candidats/lire-les-annonces.html"
	_ = m.ScrapPageNumMoniteur()

	for i := range make([]int, 10, 16) {

		if i != 0 {
			url = fmt.Sprintf("https://www.lemoniteurdespharmacies.fr/emploi/espace-candidats/lire-les-annonces-%d.html", i)
		}
		fmt.Println(url)
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
						m.Urls = append(m.Urls, node.Attributes[1])
					}
				}
				return nil
			}),
		)
		if err != nil {
			fmt.Println(err)
		}
	}

}

func (m *ScrapperPharma) ScrapPageNumMoniteur() int {

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

func (m *ScrapperPharma) ExtractDepartement(str string) int {

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

func (m *ScrapperPharma) WrapperScrappUrl() {
	m.ScrappMoniteurUrls()
	fmt.Println(m.Urls)
}
func (m *ScrapperPharma) WrapperScrappAnnonces() {
	if len(m.Urls) == 0 {
		m.ScrappMoniteurUrls()
		utils.ArrToJson(m.Urls, "moniteur_urls.json")
		fmt.Println(m.Urls)
	}
	annonces := m.ScrappAnnonces(m.Selectors)

	err := utils.ArrToJson(annonces, "moniteur.json")
	fmt.Println(err)

}
