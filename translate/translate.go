package translate

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	translate "cloud.google.com/go/translate/apiv3"
	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/aureliengasser/planetocd/articles"
	"github.com/aureliengasser/planetocd/server"
	"github.com/gomarkdown/markdown"
	translatepb "google.golang.org/genproto/googleapis/cloud/translate/v3"
)

// HTMLToMarkdown ...
func HTMLToMarkdown() {
	html, _ := ioutil.ReadFile("./translate/_input.html")
	converter := md.NewConverter("", true, nil)

	markdown, err := converter.ConvertString(string(html))
	if err != nil {
		log.Fatal(err)
	}
	ioutil.WriteFile("./translate/_input.md", []byte(markdown), 0644)
}

// CreateTranslatedArticle ....
func CreateTranslatedArticle(id string, originalURL string, originalAuthor string, originalTitle string) {
	inputFileMD := "./translate/_input.md"
	inputFileHTML := "./translate/_input.html"
	inputMD, _ := ioutil.ReadFile(inputFileMD)

	html := markdown.ToHTML(inputMD, nil, nil)

	metadata := articles.ArticleMetadata{
		OriginalURL:    originalURL,
		OriginalTitle:  originalTitle,
		OriginalAuthor: originalAuthor,
		Languages:      make(map[string]articles.ArticleLanguageMetadata),
	}

	for _, lang := range server.SupportedLanguages {
		translateAndWrite(lang, string(html), id)
		translatedTitle, err := translateText(os.Stdout, "planetocd", "en", lang, originalTitle, "text/plain", "default")
		if err != nil {
			log.Fatal(err)
		}
		metadata.Languages[lang] = articles.ArticleLanguageMetadata{
			Title: translatedTitle,
		}
	}

	metadataJSON, err := json.MarshalIndent(&metadata, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	slug := server.Slugify(originalTitle)
	ioutil.WriteFile("./articles/articles/"+id+"__"+slug+".json", metadataJSON, 0644)

	copyFile(inputFileMD, "./articles/articles/"+id+"__original.md")
	copyFile(inputFileHTML, "./articles/articles/"+id+"__original.html")
}

func translateAndWrite(lang string, html string, id string) {
	translatedHTML, err := translateText(os.Stdout, "planetocd", "en", lang, html, "text/html", "default")
	if err != nil {
		log.Fatal(err)
	}

	converter := md.NewConverter("", true, nil)
	markdown, err := converter.ConvertString(translatedHTML)
	if err != nil {
		log.Fatal(err)
	}
	ioutil.WriteFile("./articles/articles/"+id+"_"+lang+".md", []byte(markdown), 0644)
}

func copyFile(src string, dest string) {
	input, err := ioutil.ReadFile(src)
	if err != nil {
		log.Fatal(err)
		return
	}

	err = ioutil.WriteFile(dest, input, 0644)
	if err != nil {
		fmt.Println("Error creating", dest)
		log.Fatal(err)
		return
	}
}

func translateText(w io.Writer, projectID string, sourceLang string, targetLang string, text string, mimeType string, modelType string) (string, error) {
	ctx := context.Background()
	client, err := translate.NewTranslationClient(ctx)
	if err != nil {
		return "", fmt.Errorf("NewTranslationClient: %v", err)
	}
	defer client.Close()

	model := ""
	if modelType != "default" {
		model = "projects/planetocd/locations/global/models/general/" + modelType
	}

	req := &translatepb.TranslateTextRequest{
		Parent:             fmt.Sprintf("projects/%s/locations/global", projectID),
		SourceLanguageCode: sourceLang,
		TargetLanguageCode: targetLang,
		Model:              model,    // nmt or base
		MimeType:           mimeType, // Mime types: "text/plain", "text/html"
		Contents:           []string{text},
	}

	resp, err := client.TranslateText(ctx, req)
	if err != nil {
		return "", fmt.Errorf("TranslateText: %v", err)
	}

	res := ""
	for _, translation := range resp.GetTranslations() {
		res = res + translation.GetTranslatedText() + "\n"
	}

	return res, nil
}
