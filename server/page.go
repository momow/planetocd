package server

// Page ...
type Page struct {
	Constants map[string]interface{}
	Meta      *PageMeta
}

// PageMeta ...
type PageMeta struct {
	Lang                  string
	Description           string
	CanonicalURL          string
	Title                 string
	RootURL               string
	SocialImage           string
	EnableGoogleAnalytics bool
}

// T translates an input key using the Page's lang code
func (p *Page) T(key string) string {
	return Translate(p.Meta.Lang, key)
}

// URL adds the language prefix to an URL path
func (p *Page) URL(path string) string {
	return "/" + p.Meta.Lang + path
}
