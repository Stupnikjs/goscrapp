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
		fmt.Printf("node number :%d , node children count %d", i, node.ChildrenCount)	

 }
}

