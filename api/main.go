package main

import (
	"go_mission/api/app"
	"go_mission/api/config"
	"go_mission/api/handler"
	"log"
)

func main() {
	log.Fatalf("%+v", serve())
}

func serve() error {
	app, err := app.NewApp()
	if err != nil {
		return err
	}
	addr := ":" + config.Port()
	return handler.NewRouter(app).Run(addr)
}
