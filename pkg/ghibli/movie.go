package ghibli

// Movie model
type Movie struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Director    string `json:"director"`
	Producer    string `json:"producer"`
	ReleaseDate string `json:"release_date"`
	Score       string `json:"rt_score"`
	URL         string `json:"url"`
}
