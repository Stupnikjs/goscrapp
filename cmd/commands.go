package main

import (
	"fmt"

	"github.com/Stupnikjs/goscrapp/scrap"
)

var moniteurScrap = scrap.Moniteur
var ocp = scrap.Ocp

var commandsMap = map[string]func(){
	"exit":   Exit,
	"murl":   moniteurScrap.WrapperScrappUrl,
	"anmoni": moniteurScrap.WrapperScrappAnnonces,
	"ocpurl": ocp.GetOcpUrls,
	"anocp":  ocp.WrapperScrappAnnonces,
}

func CommandParser(cmd string) {
	f, ok := commandsMap[cmd]
	if ok {
		f()
	} else {
		fmt.Println("unknown command")
	}
}
