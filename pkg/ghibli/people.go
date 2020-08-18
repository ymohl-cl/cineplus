package ghibli

import (
	"strings"
)

// People model
type People struct {
	ID         string   `json:"id"`
	Name       string   `json:"name"`
	Gender     string   `json:"gender"`
	Age        string   `json:"late teens"`
	EyeColor   string   `json:"eye_color"`
	HairColor  string   `json:"hair_color"`
	FilmsURL   []string `json:"films"`
	SpeciesURL string   `json:"species"`
	URL        string   `json:"url"`
}

// IsInMovie check if the people is in the movie defined by the movie identifier
func (p People) IsInMovie(movieID string) bool {
	for _, url := range p.FilmsURL {
		if strings.Contains(url, movieID) {
			return true
		}
	}
	return false
}
