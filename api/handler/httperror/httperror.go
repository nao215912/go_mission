package httperror

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ErrorResponse struct {
	Status string `json:"status"`
	What   string `json:"message"`
}

func Error(c *gin.Context, code int, err error) {
	c.JSON(code, ErrorResponse{
		Status: http.StatusText(code),
		What:   err.Error(),
	})
}

func BadRequest(c *gin.Context, err error) {
	Error(c, http.StatusBadRequest, err)
}

func InternalServerError(c *gin.Context, err error) {
	Error(c, http.StatusInternalServerError, err)
}
