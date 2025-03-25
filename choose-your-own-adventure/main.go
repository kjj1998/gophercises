package main

import (
	"encoding/json"
	"html/template"
	"net/http"
	"os"

	models "github.com/kjj1998/gophercises/choose-your-own-adventure/models"
)

var mappings models.Story

// func frontPageHandler(w http.ResponseWriter, r *http.Request) {
// 	http.Redirect(w, r, "/intro", http.StatusTemporaryRedirect)
// }

func handler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	if path == "/" {
		http.Redirect(w, r, "/intro", http.StatusTemporaryRedirect)
		return
	}

	page := path[1:]
	t, _ := template.ParseFiles("templates/page.html")

	t.Execute(w, mappings[page])
}

func main() {
	file, err := os.Open("data/gopher.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	_ = decoder.Decode(&mappings)

	http.HandleFunc("/", handler)

	http.ListenAndServe(":8080", nil)
}
