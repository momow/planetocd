package articles

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"regexp"
	"strconv"
)

var articles = make(map[string]map[int]*Article)
var articlesRootPath = "articles/articles/"
var regexTranslationFile = regexp.MustCompile(`_([a-z]{2}).md$`)
var regexMetadataFile = regexp.MustCompile(`.*/([0-9]+)_[^/]+.json$`)

// GetArticles ...
func GetArticles() map[string]map[int]*Article {
	ensureLoaded()
	return articles
}

func ensureLoaded() {
	if len(articles) != 0 {
		return
	}
	articles, _ := filepath.Glob(articlesRootPath + "*__*.json")
	for _, metadataFile := range articles {
		loadArticle(metadataFile)
	}
}

func loadArticle(metadataFile string) {
	// Get article ID
	matches := regexMetadataFile.FindStringSubmatch(metadataFile)
	if len(matches) != 2 {
		log.Panic("Error parsing file name " + metadataFile)
	}
	id := matches[1]
	idN, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal(err)
	}

	// Load metadata
	metadataJSON, err := ioutil.ReadFile(metadataFile)
	if err != nil {
		log.Fatal(err)
	}
	var metadata ArticleMetadata
	err = json.Unmarshal(metadataJSON, &metadata)
	if err != nil {
		log.Fatal(err)
	}

	// Load each language
	files, _ := filepath.Glob(articlesRootPath + id + "_*.md")
	for _, langFile := range files {
		matches := regexTranslationFile.FindStringSubmatch(langFile)
		if len(matches) == 0 {
			continue // xxx__original.md
		}
		lang := matches[1]
		if len(lang) != 2 {
			log.Panic("Error parsing file name " + langFile)
		}
		article, err := loadArticleInLang(id, idN, lang, metadata)
		if err != nil {
			fmt.Println(err)
			continue
		}
		_, ok := articles[lang]
		if !ok {
			articles[lang] = make(map[int]*Article)
		}
		articles[lang][idN] = article
	}
}

func loadArticleInLang(id string, idN int, lang string, metadata ArticleMetadata) (*Article, error) {
	langMetadata := metadata.Languages[lang]
	filePath := articlesRootPath + id + "_" + lang + ".md"
	mdBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("Cannot read file %v: %v", filePath, err)
	}

	return &Article{
		ID:             idN,
		Lang:           lang,
		Markdown:       string(mdBytes),
		Title:          langMetadata.Title,
		OriginalURL:    metadata.OriginalURL,
		OriginalTitle:  metadata.OriginalTitle,
		OriginalAuthor: metadata.OriginalAuthor,
	}, nil
}
