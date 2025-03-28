package urlshort

import (
	"encoding/json"
	"net/http"

	"github.com/boltdb/bolt"
	"github.com/kjj1998/gophercises/url-short/database"
	"gopkg.in/yaml.v2"
)

type T struct {
	Path string `yaml:"path"`
	Url  string `yaml:"url"`
}

type J struct {
	Path string `json:"path"`
	Url  string `json:"url"`
}

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		url, ok := pathsToUrls[path]

		if !ok {
			fallback.ServeHTTP(w, r)
			return
		}

		http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	})
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//   - path: /some-path
//     url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(decoder *yaml.Decoder, fallback http.Handler) (http.HandlerFunc, error) {
	parsedYaml, err := parseYAML(decoder)
	if err != nil {
		return nil, err
	}

	pathMap := buildYamlMap(parsedYaml)

	return MapHandler(pathMap, fallback), nil
}

func parseYAML(decoder *yaml.Decoder) ([]T, error) {
	var mappings []T

	err := decoder.Decode(&mappings)

	return mappings, err
}

func buildYamlMap(mappings []T) map[string]string {
	pathsToUrl := make(map[string]string)

	for _, v := range mappings {
		pathsToUrl[v.Path] = v.Url
	}

	return pathsToUrl
}

func JSONHandler(bytes []byte, fallback http.Handler) (http.HandlerFunc, error) {
	parsedJson, err := parseJSON(bytes)
	if err != nil {
		return nil, err
	}

	pathMap := buildJsonMap(parsedJson)

	return MapHandler(pathMap, fallback), nil
}

func parseJSON(bytes []byte) ([]J, error) {
	var mappings []J

	err := json.Unmarshal(bytes, &mappings)

	return mappings, err
}

func buildJsonMap(mappings []J) map[string]string {
	pathsToUrl := make(map[string]string)

	for _, v := range mappings {
		pathsToUrl[v.Path] = v.Url
	}

	return pathsToUrl
}

func DBHandler(db *bolt.DB, fallback http.Handler) (http.HandlerFunc, error) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		url := database.SearchURL(db, path)

		if url == "" {
			fallback.ServeHTTP(w, r)
			return
		}

		http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	}), nil
}
