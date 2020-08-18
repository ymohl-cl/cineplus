package ghibli

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
