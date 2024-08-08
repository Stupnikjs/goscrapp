package scrap

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/Stupnikjs/goscrapp/data"
	"github.com/chromedp/chromedp"
)

type Selectors struct {
	EntepriseSelector string
	DateSelector      string
	LieuSelector      string
	EmploiSelector    string
	ContratSelector   string
	UrlSelector       string
}

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
	var entreprise, date, jobtype, employementType, location string
	ctx, _ := chromedp.NewContext(context.Background())
	ctx, cancel := context.WithTimeout(ctx, time.Second*30)
	defer cancel()

	err := chromedp.Run(
		ctx,
		chromedp.Navigate(url),
		chromedp.Text(s.Selectors.EntepriseSelector, &entreprise, chromedp.NodeVisible),
		chromedp.Text(s.Selectors.DateSelector, &date, chromedp.NodeVisible),
		chromedp.Text(s.Selectors.EmploiSelector, &jobtype, chromedp.NodeVisible),
		chromedp.Text(s.Selectors.ContratSelector, &employementType, chromedp.NodeVisible),
		chromedp.Text(s.Selectors.LieuSelector, &location, chromedp.NodeVisible),
	)

	if err != nil {
		fmt.Println(err)
	}

	dateStr := strings.Split(time.Now().String(), " ")
	a := data.Annonce{
		Url:         url,
		Id:          ParseWebID(url, s.Site),
		PubDate:     date,
		Ville:       ParseVille(location, s.Site),
		Departement: s.ParseDep(location),
		Lieu:        location,
		Profession:  jobtype,
		Contrat:     employementType,
		Created_at:  dateStr[0],
	}

	s.Annonces = append(s.Annonces, a)
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
	start := time.Now()
	for _, scrap := range s.Scrappers {
		scrap.UrlScrapper(&scrap)
		for _, url := range scrap.Urls {
			scrap.GetAnnonce(url)
		}
		fmt.Println(scrap.Annonces)
	}
	end := time.Now()
	fmt.Println(end.Sub(start))
}

// s *Scrapper Json()
