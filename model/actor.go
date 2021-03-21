package model

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/yeahyeahcore/KinoLab-web-app/conf"
)

type Actor struct {
	ID    int    `json:"ID,omitempty"`
	FName string `json:"FName,omitempty"`
	LName string `json:"LName,omitempty"`
}

func (ac *Actor) GetAll() map[string][]Actor {

	var actors map[string][]Actor

	r, err := http.Get(fmt.Sprintf("http://%s:%s/website/actors", conf.Api.Host, conf.Api.Port))
	if err != nil {
		fmt.Println(err)
	}
	defer r.Body.Close()

	json.NewDecoder(r.Body).Decode(&actors)

	return actors
}
