package character

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_mission/api/domain/object"
	"go_mission/api/handler/httperror"
	"net/http"
)

type ListResponse struct {
	Characters []*object.UserCharacterResponse `json:"characters"`
}

func (h *handler) List(c *gin.Context) {
	ctx := c.Request.Context()
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
	userCharacterRepository := h.app.Dao.UserCharacter()
	userCharacterResponse, err := userCharacterRepository.FindByUsername(ctx, user.Name)
	if err != nil {
		httperror.InternalServerError(c, err)
		return
	}
	c.JSON(http.StatusOK, &ListResponse{Characters: userCharacterResponse})
}
