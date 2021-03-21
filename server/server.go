package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/yeahyeahcore/KinoLab-web-app/conf"
)

//Start Запускает роутинг
func Start() {
	r := gin.New()

	r.Static("/css", "ui/assets/css")
	r.Static("/js", "ui/assets/js")
	r.Static("/img", "ui/assets/img")
	r.LoadHTMLGlob("ui/*.html")

	r.GET("/", index)
	r.GET("/film/:id", film)

	//Server run
	r.Run(fmt.Sprintf("%s:%s", conf.Server.Host, conf.Server.Port)) // listen and serve

}
