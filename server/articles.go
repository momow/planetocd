package server

import (
	"fmt"
	"html/template"
	"net/url"

	"github.com/aureliengasser/planetocd/articles"
)

var allArticles map[string]map[int]*article

type articleSummary struct {
	URL       *url.URL
	Title     string
	HTMLShort template.HTML
}

func getArticles(lang string) (map[int]*article, error) {
	ensureLoaded()
	res, ok := allArticles[lang]
	if !ok {
		return nil, fmt.Errorf("Unknown lang \"%v\"", lang)
	}
	return res, nil
}

func getArticle(lang string, id int) (*article, error) {
	byLang, err := getArticles(lang)
	if err != nil {
		return nil, err
	}
	article, ok := byLang[id]
	if !ok {
		return nil, fmt.Errorf("Unknown article id %v for lang %v", id, lang)
	}
	return article, nil
}

func ensureLoaded() {
	if allArticles != nil {
		return
	}
	allArticles = make(map[string]map[int]*article)
	all := articles.GetArticles()
	for lang, byLang := range all {
		allArticles[lang] = make(map[int]*article)
		for id, article := range byLang {
			allArticles[lang][id] = newArticle(article)
		}
	}
}
