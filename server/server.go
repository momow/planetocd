package server

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
)

var supportedLanguages = [...]string{"fr", "es", "zh"}
var router *mux.Router

// Listen ...
func Listen(scheme string, host string, port int) {
	router = mux.NewRouter().
		Schemes(scheme).
		Host(host).
		Subrouter()

	router.Path("/").HandlerFunc(handleSimplePage("index_en")).Name("index_en")
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	s := router.PathPrefix("/{language}").Subrouter()
	s.HandleFunc("/about", handleSimplePage("about")).Name("about")
	s.HandleFunc("", handleArticles)
	s.HandleFunc("/", handleArticles).Name("articles")

	log.Fatal(http.ListenAndServe(fmt.Sprint(":", port), router))
}

func handleArticles(w http.ResponseWriter, r *http.Request) {
	canonicalURL, _ := router.Get("articles").URL("language", getLang(r))
	p := getPage(r, canonicalURL)
	RenderTemplate(w, "articles", p)
}

func handleSimplePage(template string) func(w http.ResponseWriter, r *http.Request) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		canonicalURL, _ := router.Get(template).URL("language", getLang(r))
		p := getPage(r, canonicalURL)
		RenderTemplate(w, template, p)
	}
	return handler
}

func getPage(r *http.Request, canonicalURL *url.URL) *Page {
	lang := getLang(r)
	translate := func(str string) string { return Translate(lang, str) }

	return &Page{
		Lang:      lang,
		Constants: Constants,
		Meta: &PageMeta{
			Description:  "TODO",
			CanonicalURL: canonicalURL.String(),
			Title:        SiteName + " - " + translate("Articles_about_OCD"), // TODO
			RootURL:      getRootURL(lang).String(),
		},
	}
}

func getRootURL(lang string) *url.URL {
	rootURL, err := router.Get("articles").URL("language", lang)
	if err != nil {
		rootURL = &url.URL{}
	}
	return rootURL
}

func getLang(r *http.Request) string {
	vars := mux.Vars(r)
	return vars["language"]
}
