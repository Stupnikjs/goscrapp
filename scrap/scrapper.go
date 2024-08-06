package scrap

import (
	"github.com/Stupnikjs/goscrapp/data"
)

type Selectors struct {
	EntepriseSelector string
	DateSelector      string
	LieuSelector      string
	EmploiSelector    string
	ContratSelector   string
	UrlSelector       string
}

type ScrapperPharma struct {
	Selectors Selectors
	Annonces  []data.Annonce
	Urls      []string
}

type Scrapper interface {
	ScrappAnnonce(Selectors) []data.Annonce
	ScrappUrls(string)
	ParseDep(string) int
}
