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
		SelectorEvaluate: ``,
		Name:             "entreprise",
	},
	{
		SelectorEvaluate: ``,
		SelectorPath:     `//p[@class="city"]`,
		Name:             "lieu",
	},
	{ // //ion-col[contains(@class, "ion-no-padding") and contains(@class, "offerHeaderCol")]
		SelectorEvaluate: ``,
		SelectorPath:     `//span`,
		Name:             "emploi",
	},
	{
		SelectorEvaluate: `5+5`,
		SelectorPath:     `//title`,
		Name:             "contrat",
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
	ctx, cancel := context.WithTimeout(ctx, time.Second*200)
	defer cancel()

	err := chromedp.Run(
		ctx,
		chromedp.Navigate(url),
		ScrollWithChromeDP(ctx),
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
	s.Urls = href
	fmt.Println(len(href))

	return s

}

func ScrollWithChromeDP(ctx context.Context) chromedp.ActionFunc {

	return chromedp.ActionFunc(func(ctx context.Context) error {

		length := 0
		for length < 100 {
			prevlength := length
			err := chromedp.WaitVisible(".item.md.item-lines-default.item-fill-none.in-list.ion-activatable.ion-focusable.hydrated.item-label").Do(ctx)
			if err != nil {
				return err
			}
			err = chromedp.Evaluate(`
			document.querySelectorAll('.item.md.item-lines-default.item-fill-none.in-list.ion-activatable.ion-focusable.hydrated.item-label').length;
			`, &length).Do(ctx)
			if err != nil {
				return err
			}
			if prevlength == length {
				break
			}
			err = chromedp.Evaluate(fmt.Sprintf(`
			document.querySelectorAll('.item.md.item-lines-default.item-fill-none.in-list.ion-activatable.ion-focusable.hydrated.item-label')[%d].scrollIntoView();
			`, length-1), nil).Do(ctx)
			if err != nil {
				return err
			}
			/*
				err = chromedp.Sleep(time.Second * 2).Do(ctx)
				if err != nil {
					return err
				}
			*/
		}
		return nil
	})
}
