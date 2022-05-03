package user

import (
	"github.com/gin-gonic/gin"
	"go_mission/api/handler/auth"
	"go_mission/api/handler/httperror"
	"net/http"
)

type GetResponse struct {
	Name string `json:"name"`
}

func (h *handler) Get(c *gin.Context) {
	user, err := auth.UserOf(c)
	if err != nil {
		httperror.BadRequest(c, err)
	}
	c.JSON(http.StatusOK, &GetResponse{Name: user.Name})

}
