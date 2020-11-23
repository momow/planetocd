package server

import (
	"net/http"

	"golang.org/x/text/language"
)

var supportedLanguages [3]string
var langBaseToLang map[language.Base]string
var languageMatcher language.Matcher

func init() {
	frBase, _ := language.French.Base()
	esBase, _ := language.Spanish.Base()
	zhBase, _ := language.Chinese.Base()

	supportedLanguages = [3]string{
		"fr",
		"es",
		"zh",
	}

	langBaseToLang = map[language.Base]string{
		frBase: "fr",
		esBase: "es",
		zhBase: "zh",
	}

	matcherLanguages := [...]language.Tag{
		language.English,
		language.French,
		language.Spanish,
		language.Chinese,
	}

	languageMatcher = language.NewMatcher(matcherLanguages[:])
}

func inferLanguage(r *http.Request) string {
	acceptLanguage := r.Header["Accept-Language"]
	if len(acceptLanguage) > 0 {
		langs, _, err := language.ParseAcceptLanguage(acceptLanguage[0])
		if err != nil {
			return ""
		}
		tag, _, _ := languageMatcher.Match(langs...)
		base, _ := tag.Base()
		return langBaseToLang[base]
	}
	return ""
}
