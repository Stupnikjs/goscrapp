package main

import (
	"fmt"
)

var commandsMap = map[string]func(){
	"exit": Exit,
}

func CommandParser(cmd string) {
	f, ok := commandsMap[cmd]
	if ok {
		f()
	} else {
		fmt.Println("unknown command")
	}
}
