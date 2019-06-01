package main

import (
	"flag"
	"log"
)

var (
	sampleHTML = flag.String("sample_html", "./sample.html", "HTML sample file.")
)

func main() {

	flag.Parse()

	parsedHTMLTree, err := link.ParseHTML(*sampleHTML)
	if err != nil {
		log.Fatal(err)
	}

}
