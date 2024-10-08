package scrap

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Stupnikjs/goscrapp/data"
	"github.com/chromedp/chromedp"
)

type Selector struct {
	SelectorPath     string
	SelectorEvaluate string
	Name             string
	Value            string
}

type Selectors []Selector

type ScrapperSite struct {
	Site        string
	Selectors   Selectors
	UrlScrapper func(*ScrapperSite) *ScrapperSite
	Urls        []string
	Annonces    []data.Annonce
}

type Scrapper struct {
	Scrappers []ScrapperSite
	DB        *sql.DB
}

var Scr = Scrapper{
	Scrappers: []ScrapperSite{
		MoniteurScrapper,
		OcpScrapper,
	},
}

var Test = Scrapper{
	Scrappers: []ScrapperSite{
		ClubOffScrapper,
	},
}

func (s *ScrapperSite) GetAnnonce(url string) data.Annonce {

	ctx, cancel := chromedp.NewContext(context.Background())
	ctx, cancel = context.WithTimeout(ctx, time.Second*20)
	defer cancel()

	fmt.Println(url)
	text := ""
	err := chromedp.Run(
		ctx,
		chromedp.WaitReady("body", chromedp.ByQuery),
		chromedp.InnerHTML("html", &text),
		chromedp.Evaluate(`document.querySelector("html")`, &text),

	//s.SelectorProcessor(url)...,
	)
	fmt.Println("text", text)
	if err != nil {
		fmt.Println(err)
	}
	a := s.SelectorToAnnonce()
	a.Departement = s.ParseDep(a.Lieu)
	a.Url = url
	a.Id = ParseWebID(url, s.Site)
	s.Annonces = append(s.Annonces, a)
	return a
}

func (s *ScrapperSite) SelectorProcessor(url string) []chromedp.Action {
	a := []chromedp.Action{
		chromedp.Navigate(url),
	}
	for i := range s.Selectors {
		if s.Selectors[i].SelectorEvaluate == `` {
			b := chromedp.Text(s.Selectors[i].SelectorPath, &s.Selectors[i].Value, chromedp.NodeVisible)
			a = append(a, b)
		} else {
			a = append(a, chromedp.WaitVisible("h1"))
			b := chromedp.Evaluate(s.Selectors[i].SelectorEvaluate, &s.Selectors[i].Value)
			a = append(a, b)
		}
	}
	return a
}

func (s *ScrapperSite) SelectorToAnnonce() data.Annonce {
	a := data.Annonce{}
	for _, sel := range s.Selectors {
		switch sel.Name {
		case "date":
			a.PubDate = sel.Value
		case "lieu":
			a.Lieu = sel.Value
		case "emploi":
			a.Profession = sel.Value
		case "contrat":
			a.Contrat = sel.Value
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

	case "clubofficine":
		return ""
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

func (s *ScrapperSite) ParseDep(str string) int {
	if s.Site == "moniteur" {
		return ExtractDepartement(str)
	}
	split := strings.Split(str, ",")

	if len(split) < 2 {
		return 0
	}

	for dep := range data.Departements {
		if strings.Contains(dep, strings.TrimSpace(split[1])) && len(dep) == len(strings.TrimSpace(split[1])) {
			return data.Departements[dep]

		}
	}
	return 0
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

// ne marche pas
func (s *Scrapper) Json(annonces []data.Annonce) {

	bytes, err := json.Marshal(annonces)
	if err != nil {
		fmt.Println(err)
	}
	today := strings.Split(time.Now().String(), " ")[0]
	file, err := os.Create(today + ".json")
	if err != nil {
		fmt.Println(err)
	}
	file.Write(bytes)
	defer file.Close()
}
