package main

import (
	"github.com/yeahyeahcore/KinoLab-web-app/conf"
	"github.com/yeahyeahcore/KinoLab-web-app/server"
)

func main() {
	conf.Load("conf/config.json")
	server.Start()
}
