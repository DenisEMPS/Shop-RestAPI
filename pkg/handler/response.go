package handler

import (
	"school21_project1/types"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func NewErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, ErrorResponse{message})
}

type StatusResponse struct {
	Status string `json:"status"`
}

type DataResponse struct {
	Data []types.SupplierDAO `json:"data"`
}
