package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_mission/api/domain/object"
	"go_mission/api/handler/httperror"
	"net/http"
)

type UpdateRequest struct {
	Name string `json:"name"`
}

func (h *handler) Update(c *gin.Context) {
	ctx := c.Request.Context()
	req := new(UpdateRequest)
	if err := c.ShouldBindJSON(req); err != nil {
		httperror.BadRequest(c, err)
		return
	}
	token := c.Request.Header.Get("x-token")
	if token == "" {
		httperror.BadRequest(c, fmt.Errorf("incorrect Header"))
		return
	}
	entity := &object.User{
		Token: token,
		Name:  req.Name,
	}
	if err := entity.SetDecryptedName(token); err != nil {
		httperror.InternalServerError(c, err)
		return
	}
	repository := h.app.Dao.User()
	entity, err := repository.UpdateByName(ctx, entity, req.Name)
	if err != nil {
		httperror.InternalServerError(c, err)
		return
	}
	c.Writer.WriteHeader(http.StatusOK)
}
