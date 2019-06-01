package link

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"strings"

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

// ParseHTML parses HTML data and returns a parsed tree for the HTML.
func ParseHTML(r io.Reader) (*html.Node, error) {
	node, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	return node, nil
}

// GetLinkNodes returns a list of link type nodes
func GetLinkNodes(node *html.Node) []*html.Node {
	if node.Type == html.ElementNode && node.Data == "a" {
		return []*html.Node{node}
	}
	var nodes []*html.Node
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		nodes = append(nodes, GetLinkNodes(c)...)
	}
	return nodes
}

// GetLinks returns a list of Links
func GetLinks(nodes []*html.Node) []Link {
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
	link.Text = getLinkText(node)
	return link
}

func getLinkText(node *html.Node) string {
	if node.Type == html.TextNode {
		return strings.TrimSpace(node.Data)
	}
	if node.Type != html.ElementNode {
		return ""
	}
	var text string
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		text += getLinkText(c)
	}
	return text
}

// ExtractLinks does all the work.
func ExtractLinks(sampleFile string) ([]Link, error) {
	var links []Link
	data, err := ioutil.ReadFile(sampleFile)
	r := bytes.NewReader(data)
	parsedHTMLTree, err := ParseHTML(r)
	if err != nil {
		return nil, err
	}
	nodes := GetLinkNodes(parsedHTMLTree)
	links = GetLinks(nodes)
	return links, nil
}
