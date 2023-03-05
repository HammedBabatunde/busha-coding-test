package movie

type Movie struct {
	Title string `json:"title"`
	Crawl string `json:"opening_crawl"`
}

type MovieResponse struct {
	Results []Movie `json:"results"`
}
