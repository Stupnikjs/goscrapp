package scrap

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
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
	for _, url := range m.Urls {
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

		an := data.Annonce{
			Url:         url,
			PubDate:     date,
			Lieu:        location,
			Profession:  jobtype,
			Departement: extractDepartement(location),
			Contrat:     employementType,
			Created_at:  time.Now().Format("2022-02-06"),
		}
		annonces = append(annonces, an)
	}
	return annonces

}

func (m *ScrapperPharma) ScrappUrls() {
	urls := []string{}
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

func GetMoniteurUrls() {

	var nodes []*cdp.Node
	var selector string = `//ul[@class="tablelike"]//a/@href`
	var url string = "https://www.lemoniteurdespharmacies.fr/emploi/espace-candidats/lire-les-annonces.html"
	var urls = []string{}

	pageNum := ScrapPageNumMoniteur()
	for i := range make([]int, pageNum, 16) {
		if i != 0 {
			url = fmt.Sprintf("https://www.lemoniteurdespharmacies.fr/emploi/espace-candidats/lire-les-annonces-%d.html", i)
		}
		ScrappUrls(selector, url, nodes, &urls)
	}

	fmt.Println(urls)
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

func ScrappUrls(selector, url string, nodes []*cdp.Node, string *[]string) {
	panic("unimplemented")
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

func extractDepartement(str string) int {

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
