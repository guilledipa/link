package link

import (
	"bytes"
	"io/ioutil"

	"golang.org/x/net/html"
)

// Link represents a link (<a href="...">) in an HTML document.
type Link struct {
	Href string
	Text string
}

// ParseHTML parses an html file and  the parse tree for the HTML.
func ParseHTML(sampleFile string) (*html.Node, error) {
	data, err := ioutil.ReadFile(sampleFile)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	return html.Parse(bytes.NewReader(data))
}

// Returns a list of link type nodes
func linkNodes(n *html.Node) []*html.Node {
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}
	var nodes []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		nodes = append(nodes, linkNodes(c)...)
	}
	return nodes
}
