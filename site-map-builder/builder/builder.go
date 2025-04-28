package builder

import (
	"encoding/xml"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"

	"github.com/kjj1998/gophercises/html-link-parser/parser"
	"github.com/kjj1998/gophercises/site-map-builder/request"
)

type UrlSet struct {
	XMLName xml.Name `xml:"urlset"`
	Xmlns   string   `xml:"xmlns,attr"`
	Urls    []XmlUrl `xml:"url"`
}

type XmlUrl struct {
	Loc string `xml:"loc"`
}

var visited = make(map[string]struct{})
var queue []parser.Link

func getLinks(url string) []parser.Link {
	io, err := request.ReadPageHtml(url)
	if err != nil {
		log.Fatal(err)
	}

	links, err := parser.Parse(io)
	if err != nil {
		log.Fatal(err)
	}

	return links
}

func checkIfLinksVisited(queue *[]parser.Link, links []parser.Link) {
	for _, l := range links {
		if _, exists := visited[l.Href]; exists {
			continue
		} else {
			(*queue) = append((*queue), l)
		}
	}
}

func extractHostname(rawurl string) string {
	parsedUrl, err := url.Parse(rawurl)
	if err != nil {
		log.Fatal("Error parsing URL:", err)
	}

	return parsedUrl.Host
}

func createSiteMapXmlFile(hostname string) {
	file, err := os.Create("sitemap.xml")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close() // Important: Close the file when done

	urlSet := UrlSet{
		Xmlns: "http://www.sitemaps.org/schemas/sitemap/0.9",
	}

	for k := range visited {
		var fullUrl string

		if !strings.Contains(k, hostname) {
			fullUrl = "https://" + hostname + k
		} else {
			fullUrl = k
		}

		urlSet.Urls = append(urlSet.Urls, XmlUrl{Loc: fullUrl})
	}

	file.WriteString(xml.Header)
	enc := xml.NewEncoder(file)
	enc.Indent("  ", "    ")

	if err := enc.Encode(urlSet); err != nil {
		fmt.Printf("error: %v\n", err)
	}
}

func Build(url *string) {
	hostname := extractHostname(*url)
	links := getLinks(*url)
	queue = links
	var nextUrl string

	for len(queue) > 0 {
		link := queue[0]
		queue = queue[1:]

		if strings.HasPrefix(link.Href, "/") {
			if link.Href == "/" {
				continue
			}
			visited[link.Href] = struct{}{}
			nextUrl = (*url)[:len(*url)-1] + link.Href
			links = getLinks(nextUrl)
			checkIfLinksVisited(&queue, links)
		} else if strings.Contains(link.Href, hostname) {
			visited[link.Href] = struct{}{}
			links = getLinks(link.Href)
			checkIfLinksVisited(&queue, links)
		}
	}

	createSiteMapXmlFile(hostname)
}
