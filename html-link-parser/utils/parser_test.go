package utils

import (
	"testing"

	"golang.org/x/net/html"
)

func TestLinkParser(t *testing.T) {
	file := ReadHtmlFile("../files/ex2.html")
	doc, _ := html.Parse(file)

	links := LinkParser(doc)

	if links[0].Href != "https://www.twitter.com/joncalhoun" {
		t.Errorf("got %s; want 'https://www.twitter.com/joncalhoun'", links[0].Href)
	}
	if links[0].Text != "Check me out on twitter" {
		t.Errorf("got %s; want 'Check me out on twitter'", links[0].Href)
	}
	if links[1].Href != "https://github.com/gophercises" {
		t.Errorf("got %s; want 'https://github.com/gophercises'", links[1].Href)
	}
	if links[1].Text != "Gophercises is on" {
		t.Errorf("got %s; want 'Gophercises is on'", links[1].Href)
	}
}
