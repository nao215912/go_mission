package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
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
		httperror.BadRequest(c, fmt.Errorf("invalid token"))
		return
	}
	userRepository := h.app.Dao.User()
	user, err := userRepository.FindByToken(ctx, token)
	if err != nil {
		httperror.BadRequest(c, fmt.Errorf("invalid token"))
	}
	repository := h.app.Dao.User()
	_, err = repository.UpdateByName(ctx, user, req.Name)
	if err != nil {
		httperror.InternalServerError(c, err)
		return
	}
	c.Writer.WriteHeader(http.StatusOK)
}
