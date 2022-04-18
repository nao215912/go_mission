package character

import (
	"github.com/gin-gonic/gin"
	"go_mission/api/app"
)

type handler struct {
	app *app.App
}

func NewRouter(r *gin.RouterGroup, app *app.App) {
	h := &handler{app: app}
	r.POST("/draw", h.Draw)
}
