package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(">: ")
	for scanner.Scan() {
		CommandParser(strings.TrimSpace(scanner.Text()))
		fmt.Print(">: ")
	}

}

func Exit() {
	os.Exit(1)
}
