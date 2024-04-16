package link

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

// Link stores the parsed data in the expected structure
type Link struct {
	Href string
	Text string
}

// Parse will return a list of Links for a given HTML file or return error if
// it can parse the file.
func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	links := getLinks(doc)
	return links, nil
}

func getLinks(n *html.Node) []Link {
	var links []Link
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {

				links = append(links, Link{
					Href: a.Val,
					Text: getText(n),
				})
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		childLinks := getLinks(c)
		links = append(links, childLinks...)
	}
	return links
}

func getText(n *html.Node) string {
	var text string
	if n.Type == html.TextNode {
		text = n.Data
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		text += getText(c)
	}
	// TODO(guilledipa): Catch cases with nested objects
	// See "testdata/ex2.html" (on Github vs onGithub)
	return strings.TrimSpace(text)
}
