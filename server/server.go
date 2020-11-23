package server

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
)

var router *mux.Router

// Listen ...
func Listen(scheme string, host string, port int) {
	router = mux.NewRouter().
		Schemes(scheme).
		Host(host).
		Subrouter()

	router.Path("/").HandlerFunc(handleEnglishIndex).Name("index_en")
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	s := router.PathPrefix("/{language}").Subrouter()
	s.HandleFunc("/about", handleAbout).Name("about")
	s.HandleFunc("", handleArticles)
	s.HandleFunc("/", handleArticles).Name("articles")

	log.Fatal(http.ListenAndServe(fmt.Sprint(":", port), router))
}

func handleEnglishIndex(w http.ResponseWriter, r *http.Request) {
	lang := getLanguage(r)
	if lang != "" {
		url, err := router.Get("articles").URL("language", lang)
		if err == nil {
			http.Redirect(w, r, url.String(), http.StatusTemporaryRedirect)
		} else {
			fmt.Printf("Error getting URL: %v\n", err)
		}
	}
	canonicalURL, _ := router.Get("index_en").URL("language", getLang(r))
	title := SiteName + " - Knowledge base about Obsessive Compulsive Disorder (OCD)"

	p := getPage(r, canonicalURL, title, "")
	RenderTemplate(w, "index_en", p)
}

func handleArticles(w http.ResponseWriter, r *http.Request) {
	lang := getLang(r)
	canonicalURL, _ := router.Get("articles").URL("language", getLang(r))
	title := SiteName + " - " + Translate(lang, "Articles_about_OCD")
	description := Translate(lang, "Home_meta")

	p := getPage(r, canonicalURL, title, description)
	RenderTemplate(w, "articles", p)
}

func handleAbout(w http.ResponseWriter, r *http.Request) {
	lang := getLang(r)
	canonicalURL, _ := router.Get("about").URL("language", getLang(r))
	title := Translate(lang, "About") + " - " + SiteName

	p := getPage(r, canonicalURL, title, "")
	RenderTemplate(w, "about", p)
}

func getPage(r *http.Request, canonicalURL *url.URL, title string, description string) *Page {
	lang := getLang(r)

	return &Page{
		Lang:      lang,
		Constants: Constants,
		Meta: &PageMeta{
			Title:        title,
			Description:  description,
			CanonicalURL: canonicalURL.String(),
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
