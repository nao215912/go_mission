package user

import (
	"github.com/gin-gonic/gin"
	"go_mission/api/domain/object"
	"go_mission/api/handler/httperror"
	"net/http"
)

type CreateRequest struct {
	Name string `json:"name"`
}

type CreateResponse struct {
	Token string `json:"token"`
}

func (h *handler) Create(c *gin.Context) {
	ctx := c.Request.Context()
	req := new(CreateRequest)
	if err := c.ShouldBindJSON(req); err != nil {
		httperror.BadRequest(c, err)
		return
	}
	entity := &object.User{
		Name: req.Name,
	}
	if err := entity.SetEncryptedToken(req.Name); err != nil {
		httperror.InternalServerError(c, err)
		return
	}

	repository := h.app.Dao.User()
	entity, err := repository.Create(ctx, entity)
	if err != nil {
		httperror.InternalServerError(c, err)
		return
	}

	c.JSON(http.StatusOK, &CreateResponse{Token: entity.Token})
}
