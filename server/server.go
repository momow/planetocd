package server

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

var rootURL string

var supportedLanguages = [...]string{"fr", "es", "zh"}

// Listen ...
func Listen(domain string, port int) {
	rootURL = domain
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(fmt.Sprint(":", port), nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	lang, path, err := extractLanguage(r.URL)
	isLanguageChooser := err != nil

	translate := func(str string) string { return Translate(lang, str) }
	p := &Page{
		Lang:      lang,
		Body:      []byte("hello"),
		Constants: Constants,
		Meta: &PageMeta{
			Description:  "TODO",
			CanonicalURL: rootURL + "/" + lang + "/" + "TODO",
			Title:        SiteName + " - " + translate("Articles_about_OCD"),
			URLPrefix:    lang + "/",
			RootURL:      "/" + lang,
		},
	}
	if isLanguageChooser {
		RenderTemplate(w, "index_en", p)
	} else if path == "about" {
		RenderTemplate(w, "about", p)
	} else {
		RenderTemplate(w, "articles", p)
	}
}

func extractLanguage(url *url.URL) (string, string, error) {
	if url.Path == "" {
		return "", "", fmt.Errorf("no lang")
	}
	parts := strings.SplitN(url.Path[1:], "/", 2)
	if len(parts) == 0 {
		return "", "", fmt.Errorf("no lang")
	}
	prefix := parts[0]
	path := "/"
	if len(parts) == 2 {
		path = parts[1]
	}
	for _, lang := range supportedLanguages {
		if prefix == lang {
			return lang, path, nil
		}
	}
	return "", path, fmt.Errorf("Unsupported language %v", prefix)
}
