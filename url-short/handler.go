package urlshort

import (
	"net/http"

	"gopkg.in/yaml.v2"
)

type T struct {
	Path string `yaml:"path"`
	Url  string `yaml:"url"`
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

	pathMap := buildMap(parsedYaml)

	return MapHandler(pathMap, fallback), nil
}

func parseYAML(decoder *yaml.Decoder) ([]T, error) {
	var mappings []T

	err := decoder.Decode(&mappings)

	return mappings, err
}

func buildMap(parsedYml []T) map[string]string {
	pathsToUrl := make(map[string]string)

	for _, v := range parsedYml {
		pathsToUrl[v.Path] = v.Url
	}

	return pathsToUrl
}
