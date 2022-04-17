package main

import (
	"github.com/gin-gonic/gin"
	"go_mission/api/app"
	"go_mission/api/config"
	"go_mission/api/handler"
	"go_mission/api/handler/health"
	"log"
)

func main() {
	r := gin.Default()
	health.NewRouter(r.Group("/v1/health"))
	log.Fatalf("%+v", r.Run())
}

func serve() error {
	app, err := app.NewApp()
	if err != nil {
		return err
	}
	addr := ":" + config.Port()
	return handler.NewRouter(app).Run(addr)
}
