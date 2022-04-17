package user

import (
	"github.com/gin-gonic/gin"
	"go_mission/api/app"
)

type handler struct {
	app *app.App
}

func NewRouter(r *gin.RouterGroup, app *app.App) {
	h := &handler{app: app}
	r.POST("/create", h.Create)
	r.GET("/get", h.Get)
	r.PUT("/update", h.Update)

}
