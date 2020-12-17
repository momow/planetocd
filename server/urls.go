package server

import (
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/mozillazg/go-unidecode"
)

// Slugify ...
func Slugify(s string) string {
	s = unidecode.Unidecode(s)
	s = strings.ToLower(s)
	var re = regexp.MustCompile("[^a-z0-9-_]+")
	s = re.ReplaceAllLiteralString(s, "-")
	re = regexp.MustCompile("-{2,}")
	s = re.ReplaceAllLiteralString(s, "-")
	re = regexp.MustCompile("(-|-)$")
	s = re.ReplaceAllLiteralString(s, "")
	return s
}

func mustGetURL(name string, lang string) *url.URL {
	res, err := router.Get(name).URL("language", lang)
	if err != nil {
		panic(err)
	}
	return res
}

func mustGetArticleURL(article *article) *url.URL {
	res, err := router.Get("article").URL(
		"language", article.Lang,
		"id", strconv.Itoa(article.ID),
		"slug", article.Slug)
	if err != nil {
		panic(err)
	}
	return res
}

func getRootURL(lang string) *url.URL {
	rootURL, err := router.Get("articles").URL("language", lang)
	if err != nil {
		rootURL = &url.URL{}
	}
	return rootURL
}
