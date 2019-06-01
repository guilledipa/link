package link

import (
	"bytes"
	"fmt"
	"io/ioutil"

	"golang.org/x/net/html"
)

// Link represents a link (<a href="...">) in an HTML document.
type Link struct {
	Href string
	Text string
}

func (l Link) String() string {
	return fmt.Sprintf("Href: %s\nText: %s\n", l.Href, l.Text)
}

// parseHTML parses an html file and  the parse tree for the HTML.
func parseHTML(sampleFile string) (*html.Node, error) {
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
func getLinkNodes(node *html.Node) []*html.Node {
	if node.Type == html.ElementNode && node.Data == "a" {
		return []*html.Node{node}
	}
	var nodes []*html.Node
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		nodes = append(nodes, getLinkNodes(c)...)
	}
	return nodes
}

func getLinks(nodes []*html.Node) []Link {
	var links []Link
	for _, node := range nodes {
		links = append(links, generateLink(node))
	}
	return links
}

func generateLink(node *html.Node) Link {
	var link Link
	for _, a := range node.Attr {
		if a.Key == "href" {
			link.Href = a.Val
			break
		}
	}
	return link
}

// ExtractLinks does all the work.
func ExtractLinks(sampleFile string) ([]Link, error) {
	var links []Link
	parsedHTMLTree, err := parseHTML(sampleFile)
	if err != nil {
		return nil, err
	}
	nodes := getLinkNodes(parsedHTMLTree)
	links = getLinks(nodes)
	return links, nil
}
