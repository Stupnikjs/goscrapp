package main

import (
	"context"
	"fmt"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

type Annonce struct {
	Url         string `json:"url"`
	Lieu        string `json:"lieu"`
	Departement string `json:"Departement"`
	Profession  string `json:"profession"`
}

func NewAnnonce(url string) {

	ctx, _ := chromedp.NewContext(context.Background())
	ctx, cancel := context.WithTimeout(ctx, time.Minute*3)
	defer cancel()

	selector := `//ul[@class='liste_champs']/*`
	nodes := []*cdp.Node{}
	err := chromedp.Run(
		ctx,
		chromedp.Navigate(url),
		chromedp.Nodes(selector, &nodes),
		chromedp.ActionFunc(func(ctx context.Context) error {
			ProcessAnnonceNodes(nodes)
			return nil
		}),
	)

	if err != nil {
		fmt.Println(err)
	}

}

func ProcessAnnonceNodes(nodes []*cdp.Node) {
	for _, node := range nodes {
		fmt.Println(node.NodeValue)
	}
}
