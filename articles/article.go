package articles

// Article ...
type Article struct {
	Lang      string
	Title     string
	HTML      string
	HTMLShort string
	Markdown  string
}

// ArticleMetadata ...
type ArticleMetadata struct {
	OriginalURL    string                             `json:"originalUrl"`
	OriginalTitle  string                             `json:"originalTitle"`
	OriginalAuthor string                             `json:"originalAuthor"`
	Languages      map[string]ArticleLanguageMetadata `json:"languages"`
}

// ArticleLanguageMetadata ...
type ArticleLanguageMetadata struct {
	Title string `json:"title"`
}
