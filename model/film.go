package model

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/yeahyeahcore/KinoLab-web-app/conf"
)

type Film struct {
	ID          int
	FilmName    string
	Description string
	FilmYear    string
	Budget      float32
	FileURL     string
	PosterURL   string
	BannerURL   string
	Authors     []string
	Actors      []string
	Genres      []string
}

func (f *Film) Get(id int) map[string]Film {
	var film map[string]Film

	r, err := http.Get(fmt.Sprintf("http://%s:%s/website/film?id=%d", conf.Api.Host, conf.Api.Port, id))
	if err != nil {
		fmt.Println(err)
	}
	defer r.Body.Close()

	json_err := json.NewDecoder(r.Body).Decode(&film)
	if json_err != nil {
		fmt.Println(json_err)
	}

	return film
}
