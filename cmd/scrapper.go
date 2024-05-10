package main

import (
	"fmt"
	"io"

	"github.com/Stupnikjs/goscrapper/pkg/data"
	"github.com/chromedp/cdproto/cdp"
)

// NOT WORKING

func ProcessNode(w io.Writer, nodes []*cdp.Node) {

	for i, node := range nodes {
		fmt.Println("node number :", i)	

 }
}

