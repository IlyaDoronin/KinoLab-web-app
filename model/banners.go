package model

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/yeahyeahcore/KinoLab-web-app/conf"
)

type Banner struct {
	FilmName  string
	BannerURL string
}

func (ac *Banner) GetAll() map[string][]Banner {

	var banners map[string][]Banner

	r, err := http.Get(fmt.Sprintf("http://%s:%s/website/banners", conf.Api.Host, conf.Api.Port))
	if err != nil {
		fmt.Println(err)
	}
	defer r.Body.Close()

	json.NewDecoder(r.Body).Decode(&banners)

	return banners
}
