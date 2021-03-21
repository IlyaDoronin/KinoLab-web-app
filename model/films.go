package model

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/yeahyeahcore/KinoLab-web-app/conf"
)

//Film структура описывающая таблицу film в БД
type Films struct {
	ID        int
	FilmName  string
	PosterURL string
	Genres    []string
}

type Year struct {
	ID       int
	FilmYear int
}

func (f *Films) GetAllYears() map[string][]Year {
	var years map[string][]Year

	r, err := http.Get(fmt.Sprintf("http://%s:%s/website/years", conf.Api.Host, conf.Api.Port))
	if err != nil {
		fmt.Println(err)
	}
	defer r.Body.Close()

	json_err := json.NewDecoder(r.Body).Decode(&years)
	if json_err != nil {
		fmt.Println(json_err)
	}

	return years
}

func (f *Films) Get() map[string][]Films {
	var films map[string][]Films

	r, err := http.Get("http://localhost:5500/website/films?page=1")
	if err != nil {
		fmt.Println(err)
	}
	defer r.Body.Close()

	json_err := json.NewDecoder(r.Body).Decode(&films)
	if json_err != nil {
		fmt.Println(json_err)
	}

	return films
}
