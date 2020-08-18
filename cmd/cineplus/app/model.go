package app

// Data json model to return full data
type Data struct {
	Movies  []Movie  `json:"movies"`
	Peoples []People `json:"peoples"`
}

// Movie json description
type Movie struct {
	ID          string   `json:"identifier"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Director    string   `json:"director"`
	Producer    string   `json:"producer"`
	ReleaseDate string   `json:"release_date"`
	Score       string   `json:"score"`
	PeopleIDS   []string `json:"peopleIds"`
}

// People json description
type People struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Age    string `json:"age"`
}

// Pong model to describe the response http to a ping method
type Pong struct {
	Pong bool `json:"pong"`
}
