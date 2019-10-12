package urlshort

import (
	"net/http"

	yaml "gopkg.in/yaml.v2"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if pathsToUrls[r.URL.Path] != "" {
			http.Redirect(w, r, pathsToUrls[r.URL.Path] , 301)
		} else {
			fallback.ServeHTTP(w, r)
		}
    }
}

type redirect struct {
	path string `yaml:"path"`
	url string	`yaml:"url"`
}

func YAMLHandler(yaml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	parsedYaml, err := parseYAML(yaml)
	if err != nil {
	  return nil, err
	}
	pathMap := buildMap(parsedYaml)
	return MapHandler(pathMap, fallback), nil
}

func parseYAML(yml []byte)([]redirect, error){
	var parsed []redirect
	err := yaml.Unmarshal(yml, &parsed)
	return parsed, err
}

func buildMap(list []redirect)(map [string]string) {
	redirectMap := make(map [string]string)
	for _, red := range list {
		redirectMap[red.path] = red.url
	}
	return redirectMap
}