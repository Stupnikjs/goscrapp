package main

import (
	"fmt"
	"io"

	"github.com/chromedp/cdproto/cdp"
)

func ProcessNode(w io.Writer, nodes []*cdp.Node, obj *map[string]string) {

	for _, node := range nodes {
		obj := *obj
		if node.NodeName == "#text" && node.Parent.AttributeValue("class") == "Cuprum" {
			fmt.Fprintln(w, node.NodeValue)
			obj["name"] = node.NodeValue

		}

		if node.NodeName == "#text" && node.Parent.Parent.AttributeValue("class") == "col-md-12 textleft" && node.Parent.NodeName == "P" {
			obj["city"] = node.NodeValue

		}
		if node.Parent.NodeName == "TIME" {
			if node.NodeName == "SPAN" {
				obj["day"] = node.NodeValue
			}
			if node.NodeName == "SPAN" {
				obj["day"] = node.NodeValue
			}
			if node.NodeName == "SPAN" {
				obj["day"] = node.NodeValue
			}

		}

		if node.AttributeValue("class") == "panel-container event-mosaic-bg" {
			runid := node.AttributeValue("id")[3:]
			obj["runid"] = runid

		}

		if node.ChildNodeCount > 0 {
			ProcessNode(w, node.Children, &obj)
		}
		if len(obj) > 4 {
			fmt.Print("/n", obj)
		}

		obj = map[string]string{}

	}
}
