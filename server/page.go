package server

import "net/url"

type page struct {
	Constants map[string]interface{}
	Body      interface{}
	Meta      *pageMeta
}

type pageMeta struct {
	Lang                  string
	Description           string
	CanonicalURL          string
	SocialURL             string
	Title                 string
	RootURL               string
	SocialImage           string
	EnableGoogleAnalytics bool
	DisableHeaderLinks    bool
}

// T translates an input key using the Page's lang code
func (p *page) T(key string) string {
	return Translate(p.Meta.Lang, key)
}

// URL adds the language prefix to an URL path
func (p *page) URL(path string) string {
	return "/" + p.Meta.Lang + path
}

// MustGetURL returns the URL for a name, and panics if not found
func (p *page) MustGetURL(name string) *url.URL {
	return mustGetURL(name, p.Meta.Lang)
}
