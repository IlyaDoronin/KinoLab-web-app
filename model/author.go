package model

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/yeahyeahcore/KinoLab-web-app/conf"
)

//Author структура описывающая таблицу author в БД
type Author struct {
	ID    int
	FName string
	LName string
}

func (ac *Author) GetAll() map[string][]Author {

	var authors map[string][]Author

	r, err := http.Get(fmt.Sprintf("http://%s:%s/website/authors", conf.Api.Host, conf.Api.Port))
	if err != nil {
		fmt.Println(err)
	}
	defer r.Body.Close()

	json.NewDecoder(r.Body).Decode(&authors)

	return authors
}
