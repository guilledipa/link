package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/guilledipa/link"
)

func main() {
	filename := flag.String("file", "", "The HTML file to get links from.")
	flag.Parse()
	f, err := os.Open(*filename)
	if err != nil {
		log.Fatal(err)
	}
	links, err := link.Parse(f)
	if err != nil {
		log.Fatal(err)
	}
	for _, l := range links {
		fmt.Printf("Href: %s\nText: %s\n---\n", l.Href, l.Text)
	}
}
