package scrap

import (
	"github.com/Stupnikjs/goscrapp/data"
)

type Selectors struct {
	DateSelector    string
	LieuSelector    string
	EmploiSelector  string
	ContratSelector string
	UrlSelector     string
}

type ScrapperStruct struct {
	Selectors Selectors
	Annonces  []data.Annonce
	Urls      []string
}

type Scrapper interface {
	ScrappAnnonce()
	ScrappUrls()
	ParseDep()
}
