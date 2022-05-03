package gacha

import (
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

	characters, err := h.app.Dao.Character().FindByRand(ctx, req.Times)
	if err != nil {
		httperror.InternalServerError(c, err)
		return
	}

	_, err = h.app.Dao.UserCharacter().Create(ctx, userCharacters(user, characters, req))
	if err != nil {
		httperror.InternalServerError(c, err)
		return
	}

	c.JSON(http.StatusOK, &DrawResponse{Results: characters})
}

func userCharacters(user *object.User, characters []*object.Character, req *DrawRequest) []*object.UserCharacter {
	userCharacters := make([]*object.UserCharacter, 0, req.Times)
	for i := 0; i < req.Times; i++ {
		userCharacter := &object.UserCharacter{
			UserID:      user.ID,
			CharacterID: characters[i].ID,
		}
		userCharacters = append(userCharacters, userCharacter)
	}
	return userCharacters
}
