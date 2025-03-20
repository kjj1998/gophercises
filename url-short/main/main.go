package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	urlshort "github.com/kjj1998/gophercises/url-short"
	"gopkg.in/yaml.v2"
)

type T struct {
	Path string `yaml:"path"`
	Url  string `yaml:"url"`
}

func main() {
	yamlFile := flag.String("yaml", "paths.yaml", "a yaml file containing paths and the urls each of them redirect to")
	flag.Parse()

	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the fallback
	file, err := os.Open(*yamlFile)
	if err != nil {
		panic(err)
	}

	decoder := yaml.NewDecoder(file)

	yamlHandler, err := urlshort.YAMLHandler(decoder, mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
