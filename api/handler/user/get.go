package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_mission/api/handler/httperror"
	"net/http"
)

type GetResponse struct {
	Name string `json:"name"`
}

func (h *handler) Get(c *gin.Context) {
	ctx := c.Request.Context()
	token := c.Request.Header.Get("x-token")
	if token == "" {
		httperror.BadRequest(c, fmt.Errorf("incorrect Header"))
		return
	}
	userRepository := h.app.Dao.User()
	user, err := userRepository.FindByToken(ctx, token)
	if err != nil {
		httperror.BadRequest(c, fmt.Errorf("invalid token"))
	}
	c.JSON(http.StatusOK, &GetResponse{Name: user.Name})

}
