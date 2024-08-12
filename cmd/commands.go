package main

import (
	"fmt"

	"github.com/Stupnikjs/goscrapp/scrap"
)

var commandsMap = map[string]func(){
	"exit":  Exit,
	"scrap": scrap.Scr.Wrapper,
	"an":    scrap.Scr.PrintAnnnonces,
	"json":  scrap.Scr.Json,
}

func CommandParser(cmd string) {
	f, ok := commandsMap[cmd]
	if ok {
		f()
	} else {
		fmt.Println("unknown command")
	}
}
