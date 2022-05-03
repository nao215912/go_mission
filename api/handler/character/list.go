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
	user, err := auth.UserOf(c)
	if err != nil {
		httperror.BadRequest(c, err)
	}

	userCharacterResponse, err := h.app.Dao.UserCharacter().FindByUsername(c.Request.Context(), user.Name)
	if err != nil {
		httperror.InternalServerError(c, err)
		return
	}

	c.JSON(http.StatusOK, &ListResponse{Characters: userCharacterResponse})
}
