package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"school21_project1/types"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createProduct(c *gin.Context) {
	var product types.Product

	if err := c.BindJSON(&product); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Product.Create(product)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) updateProduct(c *gin.Context) {
	var productU types.ProductUpdate

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid request params")
		return
	}
	err = c.BindJSON(&productU)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Product.Update(id, productU)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, StatusResponse{
		Status: "ok",
	})
}

func (h *Handler) getProductByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid request params")
	}

	product, image, err := h.services.Product.GetByID(id)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)

	productBytes, err := json.Marshal(product)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	productPart, err := writer.CreateFormFile(fmt.Sprintf("data_product_%s", product.Name), fmt.Sprintf("data_product_%s.json", product.Name))
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	_, err = productPart.Write(productBytes)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	imagePart, err := writer.CreateFormFile(fmt.Sprintf("image_product_%s", product.Name), fmt.Sprintf("image_product_%s.png", product.Name))
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	_, err = imagePart.Write(image.Image)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	writer.Close()

	c.Header("Content-type", writer.FormDataContentType())
	c.Data(http.StatusOK, "multipart/mixed", buffer.Bytes())
}

func (h *Handler) getAllProducts(c *gin.Context) {

	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid request params")
		return
	}

	offset, err := strconv.Atoi(c.Query("offset"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid request params")
		return
	}

	products, images, err := h.services.Product.GetAll(offset, limit)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)

	for i, data := range products {
		dataBytes, err := json.Marshal(data)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal JSON"})
			return
		}

		dataPart, err := writer.CreateFormFile(fmt.Sprintf("data_product_%d", i), fmt.Sprintf("data_product_%d.json", i))
		if err != nil {
			NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		_, err = dataPart.Write(dataBytes)
		if err != nil {
			NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		imagePart, err := writer.CreateFormFile(fmt.Sprintf("image_product_%d", i), fmt.Sprintf("image_product_%d.png", i))
		if err != nil {
			NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		_, err = imagePart.Write(images[i].Image)
		if err != nil {
			NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}

	err = writer.Close()
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Header("Content-type", writer.FormDataContentType())
	c.Data(http.StatusOK, "multipart/mixed", buffer.Bytes())
}

func (h *Handler) deleteProductByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid request params")
		return
	}

	err = h.services.Product.Delete(id)

	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, StatusResponse{
		Status: "ok",
	})
}
