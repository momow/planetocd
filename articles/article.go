package articles

// Article ...
type Article struct {
	ID             int
	Lang           string
	Title          string
	Markdown       string
	OriginalURL    string
	OriginalAuthor string
	OriginalTitle  string
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
