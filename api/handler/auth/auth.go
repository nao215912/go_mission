package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_mission/api/app"
	"go_mission/api/domain/object"
)

func Middleware(app *app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		token := c.Request.Header.Get("x-token")
		if token == "" {
			return
		}
		if user, err := app.Dao.User().FindByToken(ctx, token); err != nil {
			return
		} else if user == nil {
			return
		} else {
			c.Set("x-token", user)
		}
	}
}

func UserOf(c *gin.Context) (*object.User, error) {
	if value, exists := c.Get("x-token"); exists {
		if user, ok := value.(*object.User); ok {
			return user, nil
		}
	}
	return nil, fmt.Errorf("invalid token")
}
