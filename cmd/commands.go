package main

import (
	"fmt"

	"github.com/Stupnikjs/goscrapp/data"
	"github.com/Stupnikjs/goscrapp/scrap"
)

var commandsMap = map[string]func(){
	"exit":   Exit,
	"anmoni": scrap.CreateMoniteurAnnoncesFile,
	"anocp":  scrap.CreateOcpAnnoncesFile,
	"murl":   scrap.GetMoniteurUrls,
	"ourl":   scrap.GetOcpUrls,
	"melt":   data.MeltJsonAnnonces,
}

func CommandParser(cmd string) {
	f, ok := commandsMap[cmd]
	if ok {
		f()
	} else {
		fmt.Println("unknown command")
	}
}
