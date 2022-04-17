package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_mission/api/domain/object"
	"go_mission/api/handler/httperror"
	"net/http"
)

type GetResponse struct {
	Name string `json:"name"`
}

func (h *handler) Get(c *gin.Context) {
	token := c.Request.Header.Get("x-token")
	if token == "" {
		httperror.BadRequest(c, fmt.Errorf("incorrect Header"))
		return
	}
	entity := &object.User{
		Token: token,
	}
	if err := entity.SetDecryptedName(token); err != nil {
		httperror.InternalServerError(c, err)
		return
	}
	c.JSON(http.StatusOK, &GetResponse{Name: entity.Name})

}
