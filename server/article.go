package server

import (
	"html/template"
	"net/url"

	"github.com/aureliengasser/planetocd/articles"
)

type article struct {
	*articles.Article
	HTML      template.HTML
	HTMLShort template.HTML
	Slug      string
	URL       *url.URL
}

func newArticle(a *articles.Article) *article {
	res := &article{
		Article:   a,
		HTML:      template.HTML(a.HTML),
		HTMLShort: template.HTML(a.HTMLShort),
		Slug:      Slugify(a.Title),
	}

	res.URL = mustGetArticleURL(res)
	return res
}
