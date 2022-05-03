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
	req := new(CreateRequest)
	if err := c.ShouldBindJSON(req); err != nil {
		httperror.BadRequest(c, err)
		return
	}

	user := &object.User{
		Name: req.Name,
	}
	user, err := h.app.Dao.User().Create(c.Request.Context(), user)
	if err != nil {
		httperror.InternalServerError(c, err)
		return
	}

	c.JSON(http.StatusOK, &CreateResponse{Token: user.Token})
}
