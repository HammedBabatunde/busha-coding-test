package movie

import (
	"strings"
)

type Movie struct {
	Title      string `json:"title"`
	Crawl      string `json:"opening_crawl"`
	RelaseDate string `json:"release_date"`
}

func (m *Movie) FormatCrawl() {
	m.Crawl = strings.ReplaceAll(m.Crawl, "\r\n", " ")
}

type ServerResponseMovieType struct {
	Movie
	CommentCount int64 `json:"comment_count"`
}

type MovieResponse struct {
	Results []Movie `json:"results"`
}
