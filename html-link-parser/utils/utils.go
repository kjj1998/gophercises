package utils

import (
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type Link struct {
	Href string
	Text string
}

func (l Link) String() string {
	return fmt.Sprintf(`Link{
  Href: "%s",
  Text: "%s",
}`, l.Href, l.Text)
}

func ReadHtmlFile(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	return file
}

func LinkParser(doc *html.Node) []Link {
	var links []Link

	for n := range doc.Descendants() {
		if n.Type == html.ElementNode && n.DataAtom == atom.A {
			for _, a := range n.Attr {
				if a.Key == "href" {
					link := Link{
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
