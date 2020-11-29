package server

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

var translations = make(map[string]map[string]string)

// Translate ...
func Translate(lang string, key string) string {
	return translations[lang][key]
}

func loadTranslations(lang string) map[string]string {
	var trans map[string]string
	yamlFile, err := ioutil.ReadFile("server/translations/" + lang + ".yaml")
	if err != nil {
		panic(fmt.Sprintf("Error opening translations file: %v ", err))
	}
	err = yaml.Unmarshal(yamlFile, &trans)
	if err != nil {
		panic(fmt.Sprintf("Error loading translations: %v ", err))
	}

	return trans
}

func init() {
	for _, lang := range SupportedLanguages {
		translations[lang] = loadTranslations(lang)
	}
}
