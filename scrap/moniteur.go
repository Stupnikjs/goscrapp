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

// creer un checkeur de doublons

var Moniteur = ScrapperPharma{}

var MoniteurSelectors = Selectors{
	EntepriseSelector: `//*[@itemprop='hiringOrganization']//span[@itemprop="name"]`,
	DateSelector:      `//span[@itemprop='datePosted']`,
	EmploiSelector:    `//span[@itemprop='occupationalCategory']`,
	ContratSelector:   `//span[@itemprop='employmentType']`,
	LieuSelector:      `//span[@itemprop='jobLocation']//span`,
}

func (m *ScrapperPharma) ScrappAnnonces(sels Selectors) []data.Annonce {
	annonces := []data.Annonce{}
	for _, url := range m.Urls[:5] {
		var entreprise, date, jobtype, employementType, location string

		ctx, _ := chromedp.NewContext(context.Background())
		ctx, cancel := context.WithTimeout(ctx, time.Second*7)
		defer cancel()

		err := chromedp.Run(
			ctx,
			chromedp.Navigate(url),
			chromedp.Text(sels.EntepriseSelector, &entreprise, chromedp.NodeVisible),
			chromedp.Text(sels.DateSelector, &date, chromedp.NodeVisible),
			chromedp.Text(sels.EmploiSelector, &jobtype, chromedp.NodeVisible),
			chromedp.Text(sels.ContratSelector, &employementType, chromedp.NodeVisible),
			chromedp.Text(sels.LieuSelector, &location, chromedp.NodeVisible),
		)

		if err != nil {
			fmt.Println(err)
		}

		a := data.Annonce{
			Url:         url,
			PubDate:     date,
			Lieu:        location,
			Profession:  jobtype,
			Departement: m.ExtractDepartement(location),
			Contrat:     employementType,
			Created_at:  time.Now().Format("2022-02-06"),
		}
		annonces = append(annonces, a)
	}
	return annonces

}

func (m *ScrapperPharma) ScrappUrls() []string {
	var selector string = `//ul[@class="tablelike"]//a/@href`
	var urls = []string{}
	var url string = "https://www.lemoniteurdespharmacies.fr/emploi/espace-candidats/lire-les-annonces.html"
	pageNum := m.ScrapPageNumMoniteur()

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
						urls = append(urls, node.Attributes[1])
					}
				}
				return nil
			}),
		)
		if err != nil {
			fmt.Println(err)
		}
	}
	return urls

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
	urls := m.ScrappUrls()
	fmt.Println(urls)

}
func (m *ScrapperPharma) WrapperScrappAnnonces() {
	annonces := m.ScrappAnnonces(m.Selectors)
	utils.arrToJson(annonces, "moniteur.json")

}
