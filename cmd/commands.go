package main

import (
	"fmt"

	"github.com/Stupnikjs/goscrapp/scrap"
)

var moniteurScrap = scrap.Moniteur

var commandsMap = map[string]func(){
	"exit":   Exit,
	"murl":   moniteurScrap.WrapperScrappUrl,
	"anmoni": moniteurScrap.WrapperScrappAnnonces,
	"ocpurl": scrap.GetOcpUrls,
}

func CommandParser(cmd string) {
	f, ok := commandsMap[cmd]
	if ok {
		f()
	} else {
		fmt.Println("unknown command")
	}
}
