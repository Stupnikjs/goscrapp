package scrap

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Stupnikjs/goscrapp/data"
	"github.com/chromedp/chromedp"
)

type SelectorPair struct {
	Selector string
	Value    string
}

type Selectors map[string]SelectorPair

type ScrapperSite struct {
	Site        string
	Selectors   Selectors
	UrlScrapper func(*ScrapperSite) *ScrapperSite
	Urls        []string
	Annonces    []data.Annonce
}

type Scrapper struct {
	Scrappers []ScrapperSite
}

var Scr = Scrapper{
	Scrappers: []ScrapperSite{
		MoniteurScrapper,
		OcpScrapper,
	},
}

func (s *ScrapperSite) GetAnnonce(url string) {
	ctx, _ := chromedp.NewContext(context.Background())
	ctx, cancel := context.WithTimeout(ctx, time.Second*30)
	defer cancel()

	err := chromedp.Run(
		ctx,
		s.SelectorProcessor(url)...,
	)

	if err != nil {
		fmt.Println(err)
	}
	
	a := s.SelectorToAnnonce()
	a.Url = url
	s.Annonces = append(s.Annonces, a)
}

func (s *ScrapperSite) SelectorProcessor(url string) []chromedp.Action {
	a := []chromedp.Action{
		chromedp.Navigate(url),
	}
	for _, v := range s.Selectors {
		b := chromedp.Text(v.Selector, &v.Value, chromedp.NodeVisible)
		a = append(a, b)
	}
	return a
}

func (s *ScrapperSite) SelectorToAnnonce() data.Annonce {
	a := data.Annonce{}
	for k, v := range s.Selectors {
		switch k {

		case "date":
			a.PubDate = v.Value
		case "lieu":
			a.Lieu = v.Value
		case "emploi":
			a.Profession = v.Value
		case "contrat":
			a.Contrat = v.Value

		}

	}
	return a
}

func ParseWebID(url string, site string) string {
	switch site {
	case "moniteur":
		firstSplit := strings.Split(url, "-")
		if len(firstSplit) < 2 {
			return ""
		}
		secSplit := strings.Split(firstSplit[len(firstSplit)-1], ".")
		return secSplit[0]
	case "ocp":
		split := strings.Split(url, "/")
		return split[len(split)-2]

	default:
		fmt.Println("error wrong site property")
		return ""
	}

}

func ParseVille(loc string, site string) string {
	switch site {
	case "moniteur":
		return ""
	case "ocp":
		return ""

	default:
		fmt.Println("error wrong site property")
		return ""
	}

}

func (s *Scrapper) Wrapper() {
	fmt.Println("Scrapping started !! ")
	start := time.Now()
	for _, scrap := range s.Scrappers {
		scrap.UrlScrapper(&scrap)
		fmt.Println("urls scrapped")
		for i, url := range scrap.Urls[:10] {
			scrap.GetAnnonce(url)
			fmt.Println(i + 1/10)
		}
		fmt.Println(scrap.Annonces)
	}
	s.Json()
	end := time.Now()
	fmt.Println(end.Sub(start))
}

func (s *Scrapper) PrintAnnnonces() {
	start := time.Now()
	annonces := s.GetAllAnnonces()
	fmt.Println(annonces)
	end := time.Now()
	fmt.Println(end.Sub(start))
}

func (s *Scrapper) GetAllAnnonces() []data.Annonce {
	start := time.Now()
	annonces := []data.Annonce{}
	for _, scrap := range s.Scrappers {
		annonces = append(annonces, scrap.Annonces...)
	}
	end := time.Now()
	fmt.Println(end.Sub(start))
	return annonces
}

func (s *Scrapper) Json() {
	annonces := s.GetAllAnnonces()
	bytes, err := json.Marshal(annonces)
	if err != nil {
		fmt.Println(err)
	}
	today := strings.Split(time.Now().String(), " ")[0]
	file, err := os.Create(today)
	if err != nil {
		fmt.Println(err)
	}
	file.Write(bytes)
	defer file.Close()
}
