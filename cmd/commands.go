package main

import (
	"fmt"
	"os"
	"time"

	"github.com/Stupnikjs/goscrapp/data"
	"github.com/Stupnikjs/goscrapp/database"
	"github.com/Stupnikjs/goscrapp/scrap"
)

type Application struct {
	DB       database.DBRepo
	Scrapper *scrap.Scrapper
}

func (app *Application) CommandParser(cmd string) {
	switch cmd {
	case "scrap":
		app.Wrapper()
	case "exit":
		os.Exit(1)
	default:
		fmt.Println("unexpected command")
	}

}

/*      Wrappers      */
func (app *Application) Wrapper() {
	fmt.Println("Scrapping started !! ")
	start := time.Now()
	annonces := []data.Annonce{}
	for _, scrap := range app.Scrapper.Scrappers {
		scrap.UrlScrapper(&scrap)
		fmt.Println("urls scrapped")
		for _, url := range scrap.Urls {
			a := scrap.GetAnnonce(url)
			err := app.DB.InsertAnnonce(a)
			if err != nil {
				fmt.Println(err)
			}

		}
		annonces = append(annonces, scrap.Annonces...)

	}
	app.Scrapper.Json(annonces)
	end := time.Now()
	fmt.Println(end.Sub(start))
}
