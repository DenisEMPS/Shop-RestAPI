package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	userCtx = "userId"
)

func GetUserID(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		NewErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return 0, errors.New("user id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		NewErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return 0, errors.New("user id not found")
	}
	return idInt, nil
}
