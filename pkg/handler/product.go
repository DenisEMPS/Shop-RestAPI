package handler

import (
	"net/http"
	"school21_project1/types"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createProduct(c *gin.Context) {
	var product types.ProductDAO

	if err := c.BindJSON(&product); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

}

func (h *Handler) updateProduct(c *gin.Context) {

}

func (h *Handler) getProductByID(c *gin.Context) {

}

func (h *Handler) getAllProducts(c *gin.Context) {

}

func (h *Handler) deleteProductByID(c *gin.Context) {

}
