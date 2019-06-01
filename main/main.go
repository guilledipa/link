package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"golang.org/x/net/html"
)

var (
	sampleHTML = flag.String("sample_html", "./sample.html", "HTML sample file.")
)

func main() {
	data, err := ioutil.ReadFile(*sampleHTML)
	if err != nil {
		log.Fatal(err)
	}

	parsedHTMLTree, err := html.Parse(bytes.NewReader(data))
	if err != nil {
		log.Fatal(err)
	}

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					fmt.Println(a.Val)
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

}
