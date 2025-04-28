package main

import (
	"flag"

	"github.com/kjj1998/gophercises/site-map-builder/builder"
)

func main() {
	url := flag.String("url", "https://www.calhoun.io/", "URL to Joe Calhoun's blog")

	builder.Build(url)
}
