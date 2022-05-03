package gacha

import (
	"github.com/gin-gonic/gin"
	"go_mission/api/app"
	"go_mission/api/handler/auth"
)

type handler struct {
	app *app.App
}

func NewRouter(r *gin.RouterGroup, app *app.App) {
	h := &handler{app: app}
	r.Use(auth.Middleware(app))
	r.POST("/draw", h.Draw)
}
