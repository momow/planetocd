package server

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/aureliengasser/planetocd/articles"
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

	// http.Error(w, "An internal error occurred", http.StatusInternalServerError)

	log.Fatal(http.ListenAndServe(fmt.Sprint(":", port), router))
}

func handleEnglishIndex(w http.ResponseWriter, r *http.Request) {
	lang := getLanguage(r)
	if lang != "" {
		url := mustGetURL("articles", lang)
		http.Redirect(w, r, url.String(), http.StatusTemporaryRedirect)
	}
	canonicalURL := mustGetURL("index_en", "")

	title := SiteName + " - Knowledge base about Obsessive Compulsive Disorder (OCD)"

	p := getPage(w, r, canonicalURL, title, "")
	p.Meta.DisableHeaderLinks = true
	RenderTemplate(w, "index_en", p)
}

func handleArticles(w http.ResponseWriter, r *http.Request) {
	lang := getLang(r)
	canonicalURL := mustGetURL("articles", lang)

	title := SiteName + " - " + Translate(lang, "Articles_about_OCD")
	description := Translate(lang, "Home_meta")

	p := getPage(w, r, canonicalURL, title, description)
	all := articles.ListArticles("fr")
	summaries := make([]articleSummary, len(all))
	i := 0
	for articleID, j := range all {
		summaries[i] = articleSummary{
			Title:     j.Title,
			HTMLShort: j.HTMLShort,
			URL:       "" + string(articleID),
		}
		i++
	}
	p.Body = summaries
	RenderTemplate(w, "articles", p)
}

func handleAbout(w http.ResponseWriter, r *http.Request) {
	lang := getLang(r)
	canonicalURL := mustGetURL("about", lang)
	title := Translate(lang, "About") + " - " + SiteName

	p := getPage(w, r, canonicalURL, title, "")
	RenderTemplate(w, "about", p)
}

func getPage(w http.ResponseWriter, r *http.Request, canonicalURL *url.URL, title string, description string) *page {
	lang := getLang(r)
	imageURL, err := router.Get("static").URL()
	if err != nil {
		panic(err)
	}

	return &page{
		Constants: Constants,
		Meta: &pageMeta{
			Lang:                  lang,
			Title:                 title,
			Description:           description,
			CanonicalURL:          canonicalURL.String(),
			RootURL:               getRootURL(lang).String(),
			SocialImage:           imageURL.String() + "images/logo_social.png", // TODO: article image
			EnableGoogleAnalytics: !isLocalEnvironment,
		},
	}
}

func getLang(r *http.Request) string {
	vars := mux.Vars(r)
	return vars["language"]
}
