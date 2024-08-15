package scrap

import (
	"context"
	"fmt"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

// creer un checkeur de doublons

var ClubOffSelectors = []Selector{
	{
		SelectorPath: `md hydrated`,
		Name:         "entreprise",
	},
	{
		SelectorPath: `//span[@itemprop='jobLocation']//span`,
		Name:         "lieu",
	},
	{
		SelectorPath: `//span[@itemprop='occupationalCategory']`,
		Name:         "emploi",
	},
	{
		SelectorPath: `//span[@itemprop='employmentType']`,
		Name:         "contrat",
	},
}

var ClubOffScrapper = ScrapperSite{
	Site:        "clubofficine",
	Selectors:   ClubOffSelectors,
	UrlScrapper: ScrappClubOffUrls,
}

func ScrappClubOffUrls(s *ScrapperSite) *ScrapperSite {
	var selector string = `//a[@class="item-native"]/@href`
	var url string = "https://www.clubofficine.fr/rechercher/offres"

	var nodes []*cdp.Node
	ctx, _ := chromedp.NewContext(context.Background())
	ctx, cancel := context.WithTimeout(ctx, time.Second*20)
	defer cancel()

	err := chromedp.Run(
		ctx,
		chromedp.Navigate(url),
		// chromedp.Evaluate(`window.scrollBy(0, window.innerHeight);`, nil),
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
	fmt.Println(s.Urls)
	return s

}
