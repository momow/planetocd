package server

import (
	"net/http"

	"golang.org/x/text/language"
)

// ...
var SupportedLanguages [2]string
var langBaseToLang map[language.Base]string
var languageMatcher language.Matcher

func init() {
	frBase, _ := language.French.Base()
	esBase, _ := language.Spanish.Base()
	// zhBase, _ := language.Chinese.Base()

	SupportedLanguages = [2]string{
		"fr",
		"es",
		// "zh",
	}

	langBaseToLang = map[language.Base]string{
		frBase: "fr",
		esBase: "es",
		// zhBase: "zh",
	}

	matcherLanguages := [...]language.Tag{
		language.English,
		language.French,
		language.Spanish,
		// language.Chinese,
	}

	languageMatcher = language.NewMatcher(matcherLanguages[:])
}

func getLanguage(r *http.Request) string {
	acceptLanguage := r.Header["Accept-Language"]
	if len(acceptLanguage) == 0 {
		return ""
	}
	for _, lang := range acceptLanguage {
		tag, _ := language.MatchStrings(languageMatcher, lang)
		base, _ := tag.Base()
		return langBaseToLang[base]
	}
	return ""
}
