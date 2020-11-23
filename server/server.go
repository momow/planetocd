package server

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
)

var router *mux.Router
var isLocalEnvironment bool

// Listen ...
func Listen(scheme string, host string, port int, isLocal bool) {
	isLocalEnvironment = isLocal

	router = mux.NewRouter().
		Schemes(scheme).
		Host(host).
		Subrouter()

	router.Path("/").HandlerFunc(handleEnglishIndex).Name("index_en")
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static")))).Name("static")

	s := router.PathPrefix("/{language}").Subrouter()
	s.HandleFunc("/about", handleAbout).Name("about")
	s.HandleFunc("", handleArticles)
	s.HandleFunc("/", handleArticles).Name("articles")

	log.Fatal(http.ListenAndServe(fmt.Sprint(":", port), router))
}

func handleEnglishIndex(w http.ResponseWriter, r *http.Request) {
	lang := getLanguage(r)
	if lang != "" {
		url, err := router.Get("articles").URL("language", getLang(r))
		if err != nil {
			internalServerError(w, r, err)
			return
		}
		http.Redirect(w, r, url.String(), http.StatusTemporaryRedirect)
	}
	canonicalURL, err := router.Get("index_en").URL("language", getLang(r))
	if err != nil {
		internalServerError(w, r, err)
		return
	}

	title := SiteName + " - Knowledge base about Obsessive Compulsive Disorder (OCD)"

	p, err := getPage(w, r, canonicalURL, title, "")
	if err != nil {
		return
	}
	RenderTemplate(w, "index_en", p)
}

func handleArticles(w http.ResponseWriter, r *http.Request) {
	lang := getLang(r)
	canonicalURL, err := router.Get("articles").URL("language", getLang(r))
	if err != nil {
		internalServerError(w, r, err)
		return
	}

	title := SiteName + " - " + Translate(lang, "Articles_about_OCD")
	description := Translate(lang, "Home_meta")

	p, err := getPage(w, r, canonicalURL, title, description)
	if err != nil {
		return
	}
	RenderTemplate(w, "articles", p)
}

func handleAbout(w http.ResponseWriter, r *http.Request) {
	lang := getLang(r)
	canonicalURL, err := router.Get("about").URL("language", getLang(r))
	if err != nil {
		internalServerError(w, r, err)
		return
	}
	title := Translate(lang, "About") + " - " + SiteName

	p, err := getPage(w, r, canonicalURL, title, "")
	if err != nil {
		internalServerError(w, r, err)
		return
	}
	RenderTemplate(w, "about", p)
}

func getPage(w http.ResponseWriter, r *http.Request, canonicalURL *url.URL, title string, description string) (*Page, error) {
	lang := getLang(r)
	imageURL, err := router.Get("static").URL()
	if err != nil {
		return nil, err
	}

	return &Page{
		Constants: Constants,
		Meta: &PageMeta{
			Lang:                  lang,
			Title:                 title,
			Description:           description,
			CanonicalURL:          canonicalURL.String(),
			RootURL:               getRootURL(lang).String(),
			SocialImage:           imageURL.String() + "images/logo_social.png", // TODO: article image
			EnableGoogleAnalytics: !isLocalEnvironment,
		},
	}, nil
}

func internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	fmt.Printf("ERROR serving %v: %v\n", r.URL, err)
	http.Error(w, "An internal error occurred", http.StatusInternalServerError)
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
