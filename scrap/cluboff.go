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
			var previousLength, currentLength int
			for {

				// Get the number of items before scrolling
				if err := chromedp.Evaluate(`
                    document.querySelectorAll('.item.md.item-lines-default.item-fill-none.in-list.ion-activatable.ion-focusable.hydrated.item-label').length;
                `, &previousLength).Do(ctx); err != nil {
					return err
				}
				fmt.Println(previousLength)
				// Scroll down by a fixed amount (e.g., 1000 pixels or adjust as needed)
				if err := chromedp.Evaluate(fmt.Sprintf(`
                    document.querySelectorAll('.item.md.item-lines-default.item-fill-none.in-list.ion-activatable.ion-focusable.hydrated.item-label')[%d].scrollIntoView() ;
                `, previousLength-2), nil).Do(ctx); err != nil {
					return err
				}

				// Wait for new content to load
				time.Sleep(2 * time.Second) // Adjust the sleep time as needed

				// Get the number of items after scrolling
				if err := chromedp.Evaluate(`
                    document.querySelectorAll('.item.md.item-lines-default.item-fill-none.in-list.ion-activatable.ion-focusable.hydrated.item-label').length;
                `, &currentLength).Do(ctx); err != nil {
					return err
				}

				// Check if new content was loaded
				if currentLength == previousLength {
					break // No new content loaded, exit the loop
				}
			}
			return nil
		}),
		/*
			chromedp.WaitVisible(".item.md.item-lines-default.item-fill-none.in-list.ion-activatable.ion-focusable.hydrated.item-label"),
			chromedp.Evaluate(`
			document.querySelectorAll('.item.md.item-lines-default.item-fill-none.in-list.ion-activatable.ion-focusable.hydrated.item-label')[13].scrollIntoView();
			`, nil),
			chromedp.Sleep(2*time.Second),
			chromedp.WaitVisible(".item.md.item-lines-default.item-fill-none.in-list.ion-activatable.ion-focusable.hydrated.item-label"),
			chromedp.Evaluate(`
			document.querySelectorAll('.item.md.item-lines-default.item-fill-none.in-list.ion-activatable.ion-focusable.hydrated.item-label')[26].scrollIntoView();
			`, nil),
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
	return s

}
