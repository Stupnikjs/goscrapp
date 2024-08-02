package scrap

import (
	"fmt"

	"github.com/Stupnikjs/goscrapp/data"
)

// get called by the command parser
func CreateMoniteurAnnoncesFile() {
	urls := data.OpenUrls("moniteururls.json")
	annonces := []data.Annonce{}
	for i, u := range urls {
		annonce := NewMoniteurAnnonce(u)
		fmt.Println(i)
		annonces = append(annonces, *annonce)
	}
	err := data.CreateJsonFileFromAnnonce("moniteur_annonces.json", annonces)
	if err != nil {
		fmt.Println(err)
	}
}

// get called by the command parser
func CreateOcpAnnoncesFile() {
	urls := data.OpenUrls("ocpurls.json")
	annonces := []data.Annonce{}
	for _, u := range urls {
		annonce := NewOcpAnnonce(u)
		annonces = append(annonces, *annonce)
	}
	err := data.CreateJsonFileFromAnnonce("ocp_annonces.json", annonces)
	if err != nil {
		fmt.Println(err)
	}

}
