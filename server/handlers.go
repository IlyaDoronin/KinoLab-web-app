package server

import (
	"strconv"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yeahyeahcore/KinoLab-web-app/model"
	"github.com/yeahyeahcore/KinoLab-web-app/conf"
)

var (
	ActorObj  model.Actor
	AuthorObj model.Author
	FilmsObj  model.Films
	FilmObj   model.Film
	GenreObj  model.Genre
	BannerObj model.Banner
)

func index(c *gin.Context) {
	actors := ActorObj.GetAll()
	authors := AuthorObj.GetAll()
	years := FilmsObj.GetAllYears()
	genres := GenreObj.GetAll()
	films := FilmsObj.Get()
	banners := BannerObj.GetAll()

	fmt.Println(conf.Api.Host)
	fmt.Println(conf.Api.Port)

	c.HTML(200, "index.html", gin.H{
		"Actors":  actors["actors"],
		"Authors": authors["authors"],
		"Years":   years["film_years"],
		"Genres":  genres["genres"],
		"Films":   films["films"],
		"Banners": banners["banners"],
		"Host":      conf.Api.Host,
		"Port":      conf.Api.Port,
	})
}

func film(c *gin.Context) {

	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(404, gin.H{"error": "неверный формат"})
		return
	}

	film := FilmObj.Get(ID)

	c.HTML(200, "film_page.html", gin.H{
		"ID":          film["film"].ID,
		"FilmName":    film["film"].FilmName,
		"Description": film["film"].Description,
		"FilmYear":    film["film"].FilmYear,
		"Budget":      film["film"].Budget,
		"FileURL":     film["film"].FileURL,
		"PosterURL":   film["film"].PosterURL,
		"BannerURL":   film["film"].BannerURL,
		"Authors":     film["film"].Authors,
		"Actors":      film["film"].Actors,
		"Genres":      film["film"].Genres,
		"Host":      conf.Api.Host,
		"Port":      conf.Api.Port,
	})
}
