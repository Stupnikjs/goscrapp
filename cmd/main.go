package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(">: ")
	for scanner.Scan() {

		CommandParser(strings.TrimSpace(scanner.Text()))

	}

}

// get urls from urls.txt file
func OpenUrls() []string {
	var urls = []string{}
	file, _ := os.Open("urls.txt")
	defer file.Close()
	bytes, _ := io.ReadAll(file)
	_ = json.Unmarshal(bytes, &urls)
	return urls
}
