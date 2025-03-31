package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/kjj1998/gophercises/html-link-parser/utils"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func main() {
	htmlFile := flag.String("html", "ex1.html", "a html file with <a> tags")
	flag.Parse()

	file := utils.ReadHtmlFile("files/" + *htmlFile)
	defer file.Close()

	doc, err := html.Parse(file)
	if err != nil {
		log.Fatal(err)
	}

	links := utils.LinkParser(doc)

	for _, v := range links {
		fmt.Println(v)
	}
}

func ReadHtmlFile(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	return file
}

func LinkParser(doc *html.Node) []utils.Link {
	var links []utils.Link

	for n := range doc.Descendants() {
		if n.Type == html.ElementNode && n.DataAtom == atom.A {
			for _, a := range n.Attr {
				if a.Key == "href" {
					link := utils.Link{
						Href: a.Val,
						Text: strings.TrimSpace(n.FirstChild.Data),
					}
					links = append(links, link)
					break
				}
			}
		}
	}

	return links
}
