package main

import (
	"fmt"

	"github.com/Stupnikjs/goscrapp/scrap"
)

var commandsMap = map[string]func(){
	"exit": Exit,
	"url":  scrap.Scr.Wrapper,
}

func CommandParser(cmd string) {
	f, ok := commandsMap[cmd]
	if ok {
		f()
	} else {
		fmt.Println("unknown command")
	}
}
