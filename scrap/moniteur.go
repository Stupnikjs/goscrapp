package scrap

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/Stupnikjs/goscrapp/data"
	"github.com/Stupnikjs/goscrapp/utils"
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
	var wg sync.WaitGroup
	annoncesChan := make(chan data.Annonce, len(m.Urls[:10]))

	for _, url := range m.Urls[:10] {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			a := m.GetAnnonce(url, sels)
			if a.Url != "" { // Only send if scraping was successful
				annoncesChan <- a
			}
		}(url)
	}
	// Close the channel once all goroutines are done
	go func() {
		wg.Wait()
		close(annoncesChan)
	}()

	// Collect results from the channel
	for a := range annoncesChan {
		annonces = append(annonces, a)
	}

	return annonces

}

func (m *ScrapperPharma) GetAnnonce(url string, sels Selectors) data.Annonce {
	var entreprise, date, jobtype, employementType, location string
	fmt.Println(url)
	ctx, _ := chromedp.NewContext(context.Background())
	ctx, cancel := context.WithTimeout(ctx, time.Second*30)
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

	dateStr := strings.Split(time.Now().String(), " ")
	a := data.Annonce{
		Url:         url,
		PubDate:     date,
		Lieu:        location,
		Profession:  jobtype,
		Departement: m.ExtractDepartement(location),
		Contrat:     employementType,
		Created_at:  dateStr[0],
	}
	return a
}

func (m *ScrapperPharma) ScrappUrls() {
	var selector string = `//ul[@class="tablelike"]//a/@href`
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
	m.ScrappUrls()
	fmt.Println(m.Urls)
}
func (m *ScrapperPharma) WrapperScrappAnnonces() {
	fmt.Println("here")
	if len(m.Urls) == 0 {
		m.ScrappUrls()
		utils.ArrToJson(m.Urls, "moniteur_urls.json")
		fmt.Println(m.Urls)
	}
	annonces := m.ScrappAnnonces(m.Selectors)

	err := utils.ArrToJson(annonces, "moniteur.json")
	fmt.Println(err)

}
