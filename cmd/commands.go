package main

import (
	"fmt"

	"github.com/Stupnikjs/goscrapp/data"
	"github.com/Stupnikjs/goscrapp/scrap"
)

var commandsMap = map[string]func(){
	"exit":     Exit,
	"anmoni":   CreateMoniteurAnnoncesFile,
	"anocp":    CreateOcpAnnoncesFile,
	"murl":     scrap.GetMoniteurUrls,
	"ourl":     scrap.GetOcpUrls,
	"dep":      data.ParseLieu,
	"melt":     MeltJsonAnnonces,
	"parsedep": data.ParseDep,
}

func CommandParser(cmd string) {
	f, ok := commandsMap[cmd]
	if ok {
		f()
	} else {
		fmt.Println("unknown command")
	}
}
