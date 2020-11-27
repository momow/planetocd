package main

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	translate "cloud.google.com/go/translate/apiv3"
	translatepb "google.golang.org/genproto/googleapis/cloud/translate/v3"
)

func translateText(w io.Writer, projectID string, sourceLang string, targetLang string, text string, mimeType string, modelType string) error {
	ctx := context.Background()
	client, err := translate.NewTranslationClient(ctx)
	if err != nil {
		return fmt.Errorf("NewTranslationClient: %v", err)
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
		return fmt.Errorf("TranslateText: %v", err)
	}

	for _, translation := range resp.GetTranslations() {
		fmt.Fprintf(w, "Translated text: %v\n", translation.GetTranslatedText())
	}

	return nil
}

func main() {
	dat, _ := ioutil.ReadFile("./input.html")
	err := translateText(os.Stdout, "planetocd", "en", "fr", string(dat), "text/html", "default")
	if err != nil {
		fmt.Println(err)
	}
}
