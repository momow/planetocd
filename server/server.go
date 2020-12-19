package server

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"

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
	s.HandleFunc("/articles/{id:[0-9]+}/{slug}", handleArticle).Name("article")

	// http.Error(w, "An internal error occurred", http.StatusInternalServerError)

	log.Fatal(http.ListenAndServe(fmt.Sprint(":", port), router))
}

func handleEnglishIndex(w http.ResponseWriter, r *http.Request) {
	// lang := getLanguage(r)
	// if lang != "" {
	// 	url := mustGetURL("articles", lang)
	// 	http.Redirect(w, r, url.String(), http.StatusTemporaryRedirect)
	// }
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
	all, err := getArticles(lang)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	summaries := make([]articleSummary, len(all))
	i := 0
	for _, article := range all {
		summaries[i] = articleSummary{
			Title:     article.Title,
			HTMLShort: article.HTMLShort,
			URL:       mustGetArticleURL(article),
		}
		i++
	}
	p.Body = summaries
	RenderTemplate(w, "articles", p)
}

func handleArticle(w http.ResponseWriter, r *http.Request) {
	lang := getLang(r)
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.NotFound(w, r)
		return
	}
	canonicalURL := mustGetURL("articles", lang)
	article, err := getArticle(lang, id)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	title := article.Title
	description := ""
	p := getPage(w, r, canonicalURL, title, description)
	p.Body = article
	RenderTemplate(w, "article", p)
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

	socialURL := *canonicalURL
	fmt.Printf("%v\n", socialURL.Host[:9])
	if socialURL.Host[:9] == "localhost" {
		socialURL.Host = "planetocd.org"
	}

	return &page{
		Constants: Constants,
		Meta: &pageMeta{
			Lang:                  lang,
			Title:                 title,
			Description:           description,
			CanonicalURL:          canonicalURL.String(),
			SocialURL:             socialURL.String(),
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
