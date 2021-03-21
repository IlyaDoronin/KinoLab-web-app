package model

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/yeahyeahcore/KinoLab-web-app/conf"
)

//Genre структура описывающая таблицу genre в БД
type Genre struct {
	ID        int
	GenreName string
}

func (ac *Genre) GetAll() map[string][]Genre {

	var genres map[string][]Genre

	r, err := http.Get(fmt.Sprintf("http://%s:%s/website/genres", conf.Api.Host, conf.Api.Port))
	if err != nil {
		fmt.Println(err)
	}
	defer r.Body.Close()

	json.NewDecoder(r.Body).Decode(&genres)

	return genres
}
