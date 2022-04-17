package handler

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go_mission/api/app"
	"time"
)

func NewRouter(app *app.App) *gin.Engine {
	r := gin.Default()

	useCROS(r)

	return r
}

func useCROS(r *gin.Engine) {
	cfg := cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authentication"},
		AllowCredentials: false,
		AllowAllOrigins:  true,
		MaxAge:           12 * time.Hour,
	}

	r.Use(cors.New(cfg))
}
