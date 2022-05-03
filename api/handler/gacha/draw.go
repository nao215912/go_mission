package gacha

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_mission/api/domain/object"
	"go_mission/api/handler/auth"
	"go_mission/api/handler/httperror"
	"net/http"
)

type DrawRequest struct {
	Times int `json:"times"`
}

type DrawResponse struct {
	Results []*object.Character
}

func (h *handler) Draw(c *gin.Context) {
	ctx := c.Request.Context()
	req := new(DrawRequest)
	if err := c.ShouldBindJSON(req); err != nil {
		httperror.BadRequest(c, err)
		return
	}
	user, err := auth.UserOf(c)
	if err != nil {
		httperror.BadRequest(c, err)
	}
	characters := make([]*object.Character, 0, req.Times)
	for i := 0; i < req.Times; i++ {
		character := &object.Character{
			Name: fmt.Sprintf("name %d", i),
		}
		characters = append(characters, character)
	}
	characterRepository := h.app.Dao.Character()
	characters, err = characterRepository.Create(ctx, characters)
	if err != nil {
		httperror.InternalServerError(c, err)
		return
	}
	userCharacters := make([]*object.UserCharacter, 0, req.Times)
	for i := 0; i < req.Times; i++ {
		userCharacter := &object.UserCharacter{
			UserID:      user.ID,
			CharacterID: characters[i].ID,
		}
		userCharacters = append(userCharacters, userCharacter)
	}
	userCharacterRepository := h.app.Dao.UserCharacter()
	userCharacters, err = userCharacterRepository.Create(ctx, userCharacters)
	if err != nil {
		httperror.InternalServerError(c, err)
		return
	}
	c.JSON(http.StatusOK, &DrawResponse{Results: characters})
}
