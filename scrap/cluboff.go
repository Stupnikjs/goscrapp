package scrap

import (
	"context"
	"fmt"
	"time"

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

	var url string = "https://www.clubofficine.fr/rechercher/offres"
	var href []string
	ctx, _ := chromedp.NewContext(context.Background())
	ctx, cancel := context.WithTimeout(ctx, time.Second*20)
	defer cancel()

	err := chromedp.Run(
		ctx,
		chromedp.Navigate(url),
		chromedp.WaitVisible(".item.md.item-lines-default.item-fill-none.in-list.ion-activatable.ion-focusable.hydrated.item-label"),
		chromedp.ActionFunc(func(ctx context.Context) error {
			n := 13
			err := chromedp.WaitVisible(".item.md.item-lines-default.item-fill-none.in-list.ion-activatable.ion-focusable.hydrated.item-label").Do(ctx)
			if err != nil {
				return err
			}
			err = chromedp.Evaluate(fmt.Sprintf(`
				document.querySelectorAll('.item.md.item-lines-default.item-fill-none.in-list.ion-activatable.ion-focusable.hydrated.item-label')[%d].scrollIntoView();
				`, n), nil).Do(ctx)
			if err != nil {
				return err
			}
			err = chromedp.Sleep(time.Second * 2).Do(ctx)
			if err != nil {
				return err
			}
			return nil
		}),
		/*

		 */
		chromedp.WaitVisible(".item.md.item-lines-default.item-fill-none.in-list.ion-activatable.ion-focusable.hydrated.item-label"),
		chromedp.Evaluate(`
		Array.from(document.querySelectorAll('.item.md.item-lines-default.item-fill-none.in-list.ion-activatable.ion-focusable.hydrated.item-label')).map(
		(el) => {
		return el.shadowRoot.querySelector('a').href
		})
		`, &href),
	)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(len(href))
	fmt.Println(href)
	return s

}
