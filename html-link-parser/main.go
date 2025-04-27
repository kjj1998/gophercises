package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/kjj1998/gophercises/html-link-parser/parser"
	"github.com/kjj1998/gophercises/html-link-parser/utils"
)

func main() {
	htmlFile := flag.String("html", "ex4.html", "a html file with <a> tags")
	flag.Parse()

	file := utils.ReadHtmlFile("files/" + *htmlFile)
	defer file.Close()

	links, err := parser.Parse(file)
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range links {
		fmt.Println(v)
	}
}
