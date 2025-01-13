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

// CreateProduct godoc
//
//	@Summary		Create Product
//	@Description	Create a new product with the provided data
//	@Tags			product
//	@Accept			json
//	@Produce		json
//	@Param			input	body		types.Product			true		"Product Info"
//	@Success		201		{object}	map[string]interface{}	"id" :int	"Successful response with product ID"
//	@Failure		400		{object}	ErrorResponse	"invalid request params"
//	@Failure		500		{object}	ErrorResponse	"internal server error"
//	@Router			/product [post]
func (h *Handler) CreateProduct(c *gin.Context) {
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

// UpdateProduct godoc
//
//	@Summary		Update Product
//	@Description	Update product quantity
//	@Tags			product
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int						true		"Product ID"
//	@Param			input	body		types.ProductUpdate		true		"Product update info"
//	@Success		200		{object}	StatusResponse	"status":	"ok"	"Successful response"
//	@Failure		400		{object}	ErrorResponse	"invalid request params"
//	@Failure		500		{object}	ErrorResponse	"internal server error"
//	@Router			/product/{id} [patch]
func (h *Handler) UpdateProduct(c *gin.Context) {
	var productU types.ProductUpdate

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid request params")
		return
	}
	err = c.BindJSON(&productU)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid request params")
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

// GetProductByID godoc
//
//	@Summary		Get Product by ID
//	@Description	This endpoint returns product and image as a multipart/mixed response. The response includes a JSON part with product and binary parts for image.
//	@Tags			product
//	@Accept			json
//	@Produce		multipart/mixed
//	@Param			id	path		int														true	"Product ID"
//	@Success		200	{file} data_product_.json image_product_.png "Multipart/mixed - json-data of Product with corresponding binary part Image"
//	@Failure		400	{object}	ErrorResponse									"invalid request params"
//	@Failure		404	{object}	ErrorResponse									"Not Found"
//	@Failure		500	{object}	ErrorResponse									"internal server error"
//	@Router			/product/{id} [get]
func (h *Handler) GetProductByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid request params")
	}

	product, image, err := h.services.Product.GetByID(id)
	if err != nil {
		NewErrorResponse(c, http.StatusNotFound, err.Error())
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

// GetAllProducts godoc
//
//	@Summary		Get All products
//	@Description	This endpoint returns products and images as a multipart/mixed response. The response includes a JSON part with products and binary parts for images.
//	@Tags			product
//	@Accept			json
//	@Produce		multipart/mixed
//	@Param			limit	query		int															false	"Limit"
//	@Param			offset	query		int															false	"Offset"
//	@Success		200		{file}	data_product_.json image_product_..png "Multipart/mixed array of json-data Products with corresponding binary part Images"
//	@Failure		400		{object}	ErrorResponse									"invalid request params"
//	@Failure		404		{object}	ErrorResponse									"Not Found"
//	@Failure		500		{object}	ErrorResponse										"internal server error"
//	@Router			/product [get]
func (h *Handler) GetAllProducts(c *gin.Context) {

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
		NewErrorResponse(c, http.StatusNotFound, err.Error())
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

// DeleteProductByID godoc
//
//	@Summary		Delete Product
//	@Description	Delete product by ID
//	@Tags			product
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int						true		"Product ID"
//	@Success		200	{object}	StatusResponse	"status":	"ok"	"Successful response"
//	@Failure		400	{object}	ErrorResponse	"invalid request params"
//	@Failure		500	{object}	ErrorResponse	"internal server error"
//	@Router			/product/{id} [delete]
func (h *Handler) DeleteProductByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid request params")
		return
	}

	err = h.services.Product.Delete(id)

	if err != nil {
		NewErrorResponse(c, http.StatusNotFound, "product was not find")
		return
	}

	c.JSON(http.StatusOK, StatusResponse{
		Status: "ok",
	})
}
