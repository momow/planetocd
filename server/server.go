package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var rootURL string

var supportedLanguages = [...]string{"fr", "es", "zh"}

// Listen ...
func Listen(domain string, port int) {
	rootURL = domain

	r := mux.NewRouter()

	r.HandleFunc("/", languageChooseHandler)

	var s *mux.Router

	for _, lang := range supportedLanguages {
		s = r.PathPrefix("/" + lang).Subrouter()
		s.HandleFunc("/about", aboutHandler(lang))
		s.HandleFunc("", homepageHandler(lang))
		s.HandleFunc("/", homepageHandler(lang))
	}

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Fatal(http.ListenAndServe(fmt.Sprint(":", port), r))
}

func aboutHandler(lang string) func(w http.ResponseWriter, r *http.Request) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		p := getPage(lang)
		RenderTemplate(w, "about", p)
	}
	return handler
}

func homepageHandler(lang string) func(w http.ResponseWriter, r *http.Request) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		p := getPage(lang)
		RenderTemplate(w, "articles", p)
	}
	return handler
}

func languageChooseHandler(w http.ResponseWriter, r *http.Request) {
	p := getPage("")
	RenderTemplate(w, "index_en", p)
}

func getPage(lang string) *Page {
	translate := func(str string) string { return Translate(lang, str) }
	return &Page{
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
}
