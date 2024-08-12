package main

import (
	"database/sql"
	"fmt"

	"github.com/Stupnikjs/goscrapp/scrap"
)

type Application struct {
	DB       *database.DBRepo
	Commands map[string]func()
}

var commandsMap = map[string]func(){
	"exit":  Exit,
	"scrap": scrap.Scr.Wrapper,
	"an":    scrap.Scr.PrintAnnnonces,
}

func (app *Application) CommandParser(cmd string) {
	f, ok := app.Commands[cmd]
	if ok {
		f()
	} else {
		fmt.Println("unknown command")
	}
}
