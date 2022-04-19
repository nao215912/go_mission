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
		httperror.BadRequest(c, fmt.Errorf("incorrect Header"))
		return
	}
	req := &object.User{
		Token: token,
	}
	if err := req.SetDecryptedName(token); err != nil {
		httperror.InternalServerError(c, err)
		return
	}
	userCharacterRepository := h.app.Dao.UserCharacter()
	userCharacterResponse, err := userCharacterRepository.FindByUsername(ctx, req.Name)
	if err != nil {
		httperror.InternalServerError(c, err)
		return
	}
	c.JSON(http.StatusOK, &ListResponse{Characters: userCharacterResponse})
}
