package handler

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go_mission/api/app"
	"go_mission/api/handler/character"
	"go_mission/api/handler/gacha"
	"go_mission/api/handler/user"
	"time"
)

func NewRouter(app *app.App) *gin.Engine {
	r := gin.Default()

	useCROS(r)

	user.NewRouter(r.Group("/user"), app)
	gacha.NewRouter(r.Group("/gacha"), app)
	character.NewRouter(r.Group("/character"), app)

	return r
}

func useCROS(r *gin.Engine) {
	cfg := cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "x-token"},
		AllowCredentials: false,
		AllowAllOrigins:  true,
		MaxAge:           12 * time.Hour,
	}

	r.Use(cors.New(cfg))
}
