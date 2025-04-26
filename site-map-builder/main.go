package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/kjj1998/gophercises/html-link-parser/utils"
	"github.com/kjj1998/gophercises/site-map-builder/request"
)

func main() {
	url := flag.String("url", "https://www.calhoun.io/", "URL to Joe Calhoun's blog")
	flag.Parse()

	doc, err := request.ReadPageHtml(*url)
	if err != nil {
		log.Fatal(err)
	}

	links := utils.LinkParser(doc)

	for _, v := range links {
		fmt.Println(v)
	}
}
