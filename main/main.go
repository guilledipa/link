package main

import (
	"flag"
	"log"

	"github.com/guilledipa/link"
)

var (
	sampleHTML = flag.String("sample_html", "./sample.html", "HTML sample file.")
)

func main() {

	flag.Parse()

	links, err := link.ExtractLinks(*sampleHTML)
	if err != nil {
		log.Fatal(err)
	}

}
