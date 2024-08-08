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
}

type ScrapperPharma struct {
	Scrappers []ScrapperSite
	Annonces  []data.Annonce
}

var Scrapper = ScrapperPharma{
	[]ScrapperSite{
		MoniteurScrapper,
		OcpScrapper,
	},
	[]data.Annonce{},
}

func (m *ScrapperPharma) GetAnnonce(url string, sels Selectors) data.Annonce {
	var entreprise, date, jobtype, employementType, location string
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
		Url:        url,
		Id:         m.ParseWebID(url),
		PubDate:    date,
		Ville:      m.ParseVille(location),
		Lieu:       location,
		Profession: jobtype,
		Contrat:    employementType,
		Created_at: dateStr[0],
	}

	if m.Selectors.Site == "moniteur" {
		a.Departement = m.ExtractDepartement(location)
	}
	if m.Selectors.Site == "ocp" {
		a.Departement = m.ParseDep(location)
	}

	return a
}

func (m *ScrapperPharma) ScrappAnnonces() []data.Annonce {
	annonces := []data.Annonce{}
	for _, url := range m.Urls {
		a := m.GetAnnonce(url, m.Selectors)
		if a.Url != "" {
			annonces = append(annonces, a)
		}
	}
	return annonces
}

func (m *ScrapperPharma) ResetUrls() {
	m.Urls = []string{}
}
func (m *ScrapperPharma) ResetAnnonces() {
	m.Annonces = []data.Annonce{}
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

func (m *ScrapperPharma) ParseVille(loc string, site string) string {
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

// iterate througth selectors

// scrap urls
