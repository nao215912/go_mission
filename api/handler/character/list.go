package character

import (
	"github.com/gin-gonic/gin"
	"go_mission/api/domain/object"
	"go_mission/api/handler/auth"
	"go_mission/api/handler/httperror"
	"net/http"
)

type ListResponse struct {
	Characters []*object.UserCharacterResponse `json:"characters"`
}

func (h *handler) List(c *gin.Context) {
	ctx := c.Request.Context()
	user, err := auth.UserOf(c)
	if err != nil {
		httperror.BadRequest(c, err)
	}
	userCharacterRepository := h.app.Dao.UserCharacter()
	userCharacterResponse, err := userCharacterRepository.FindByUsername(ctx, user.Name)
	if err != nil {
		httperror.InternalServerError(c, err)
		return
	}
	c.JSON(http.StatusOK, &ListResponse{Characters: userCharacterResponse})
}
