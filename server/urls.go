package server

import (
	"net/url"
)

func mustGetURL(name string, lang string) *url.URL {
	res, err := router.Get(name).URL("language", lang)
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
