package main

import (
	"fmt"

	"github.com/Stupnikjs/goscrapp/data"
	"github.com/Stupnikjs/goscrapp/scrap"
)

var commandsMap = map[string]func(){
	"exit":  Exit,
	"scrap": scrap.Scr.Wrapper,
	"an":    scrap.Scr.PrintAnnnonces,
}

func (app *data.Application) CommandParser(cmd string) {
	f, ok := app.Commands[cmd]
	if ok {
		f()
	} else {
		fmt.Println("unknown command")
	}
}
