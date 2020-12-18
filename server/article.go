package server

import (
	"html/template"
	"net/url"

	"github.com/aureliengasser/planetocd/articles"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
)

type article struct {
	*articles.Article
	HTML      template.HTML
	HTMLShort template.HTML
	Slug      string
	URL       *url.URL
}

func newArticle(a *articles.Article) *article {

	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	HTMLBytes := markdown.ToHTML([]byte(a.Markdown), nil, renderer)
	HTML := string(HTMLBytes)
	HTMLShort := getHTMLShort(HTML)

	res := &article{
		Article:   a,
		HTML:      template.HTML(HTML),
		HTMLShort: template.HTML(HTMLShort),
		Slug:      Slugify(a.Title),
	}

	res.URL = mustGetArticleURL(res)
	return res
}

func getHTMLShort(HTML string) string {
	endTag := "</p>"
	length := len(HTML)
	i := 500
	for ; i+len(endTag) < length; i++ {
		if HTML[i:i+len(endTag)] == endTag {
			break
		}
	}
	return HTML[:i+len(endTag)]
}
