package server

// Page ...
type Page struct {
	Lang      string
	Body      []byte
	Constants map[string]interface{}
	Meta      *PageMeta
}

// PageMeta ...
type PageMeta struct {
	Description  string
	CanonicalURL string
	Title        string
	RootURL      string
	URLPrefix    string
}

// T translates an input key using the Page's lang code
func (p *Page) T(key string) string {
	return Translate(p.Lang, key)
}

// URL adds the language prefix to an URL path
func (p *Page) URL(path string) string {
	return "/" + p.Lang + path
}
