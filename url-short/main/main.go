package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	urlshort "github.com/kjj1998/gophercises/url-short"
	database "github.com/kjj1998/gophercises/url-short/database"
	"gopkg.in/yaml.v2"
)

type T struct {
	Path string `yaml:"path"`
	Url  string `yaml:"url"`
}

func main() {
	// Open BoltDB
	db, err := database.OpenDB()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// Create bucket and populate
	database.CreateBucket(db)

	mux := defaultMux()
	yamlFile := flag.String("yaml", "paths.yaml", "a yaml file containing paths and the urls each of them redirect to")
	flag.Parse()

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

	// Build the JSONHandler using the mapHandler as the fallback
	json := `[
		{"path": "/urlshort-godoc", "url": "https://godoc.org/github.com/gophercises/urlshort"},
		{"path": "/yaml-godoc", 	"url": "https://godoc.org/gopkg.in/yaml.v2"}
	]`
	jsonHandler, err := urlshort.JSONHandler([]byte(json), mapHandler)
	if err != nil {
		panic(err)
	}

	dbHandler, err := urlshort.DBHandler(db, mux)
	if err != nil {
		panic(err)
	}

	// Start multiple servers concurrently
	go func() {
		fmt.Println("Starting JSON server on :8080")
		if err := http.ListenAndServe(":8080", jsonHandler); err != nil {
			fmt.Println("Error starting JSON server:", err)
		}
	}()

	go func() {
		fmt.Println("Starting YAML server on :8081")
		if err := http.ListenAndServe(":8081", yamlHandler); err != nil {
			fmt.Println("Error starting YAML server:", err)
		}
	}()

	go func() {
		fmt.Println("Starting DB server on :8082")
		if err := http.ListenAndServe(":8082", dbHandler); err != nil {
			fmt.Println("Error starting DB server:", err)
		}
	}()

	select {}
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
