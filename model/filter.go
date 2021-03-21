package model

//Filter структура для фильтрации информации
type Filter struct {
	Genres  []string `json:"genres,omitempty"`
	Actors  []string `json:"actors,omitempty"`
	Authors []string `json:"authors,omitempty"`
	Years   []string `json:"years,omitempty"`
	Search  string   `json:"search,omitempty"`
}

type FilterResult struct {
	ID        int
	FilmName  string
	PosterURL string
	Genres    []string
}
