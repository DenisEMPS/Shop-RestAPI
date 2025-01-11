package handler

import (
	"net/http"
	"school21_project1/types"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateImageProduct godoc
// @Summary Create Product Image
// @Description Create a new image for product
// @Tags image
// @Accept  json
// @Produce  json
// @Param input body types.CreateImageProduct true "Image and Product info"
// @Success 201 {object} map[string]interface{} "uuid":string "Successful response with image ID"
// @Failure 400 {object} response.ErrorResponse "invalid request params"
// @Failure 500 {object} response.ErrorResponse "internal server error"
// @Router /api/v1/image [post]

func (h *Handler) CreateImageProduct(c *gin.Context) {
	var image types.CreateImageProduct

	if err := c.BindJSON(&image); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid request params")
		return
	}

	uuid, err := h.services.Image.Create(image)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"uuid": uuid,
	})
}

// UpdateImageByID godoc
// @Summary Update Image
// @Description Update image by image id/uuid
// @Tags image
// @Accept  json
// @Produce  json
// @Param id path string true "Image ID"
// @Param input body types.Image true "Image bytes"
// @Success 200 {object} response.StatusResponse "status": "ok" "Successful response"
// @Failure 400 {object} response.ErrorResponse "invalid request params"
// @Failure 500 {object} response.ErrorResponse "internal server error"
// @Router /api/v1/image/{id} [patch]

func (h *Handler) UpdateImageByID(c *gin.Context) {
	uuid := c.Param("id")

	if uuid == "" {
		NewErrorResponse(c, http.StatusBadRequest, "invalid request params")
		return
	}

	var image types.Image

	err := c.BindJSON(&image)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid request params")
		return
	}

	if len(image.Image) == 0 {
		NewErrorResponse(c, http.StatusBadRequest, "invalid request params")
		return
	}

	err = h.services.Image.Update(uuid, image)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, StatusResponse{
		Status: "ok",
	})
}

// DeleteImageByID godoc
// @Summary Delete Image by ID
// @Description Delete image by image ID/uuid
// @Tags image
// @Accept  json
// @Produce  json
// @Param id path string true "Image ID"
// @Success 200 {object} map[string]interface{} "status": "ok" "Successful response"
// @Failure 400 {object} response.ErrorResponse "invalid request params"
// @Failure 500 {object} response.ErrorResponse "internal server error"
// @Router /api/v1/image/{id} [delete]

func (h *Handler) DeleteImageByID(c *gin.Context) {
	uuid := c.Param("id")

	if uuid == "" {
		NewErrorResponse(c, http.StatusBadRequest, "invalid request params")
		return
	}

	err := h.services.Image.Delete(uuid)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, StatusResponse{
		Status: "ok",
	})
}

// GetImageByProductID godoc
// @Summary Get Image by Product ID
// @Description Get image by product ID
// @Tags image
// @Accept  json
// @Produce  json
// @Param id path int true "Product ID"
// @Success 200 {object} types.Image.Image "Successful response with image"
// @Failure 400 {object} response.ErrorResponse "invalid request params"
// @Failure 500 {object} response.ErrorResponse "internal server error"
// @Router /api/v1/image/product_id/{id} [get]

func (h *Handler) GetImageByProductID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid request params")
		return
	}

	image, err := h.services.Image.GetByProductID(id)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Header("Content-Disposition", "attachment; filename=image.png")
	c.Data(http.StatusOK, "application/octet-stream", image.Image)
}

// GetImageByID godoc
// @Summary Get Image by ID
// @Description Get image by image ID/uuid
// @Tags image
// @Accept  json
// @Produce  json
// @Param id path string true "Image_id"
// @Success 200 {object} types.Image.Image "Successful response with image"
// @Failure 400 {object} response.ErrorResponse "invalid request params"
// @Failure 500 {object} response.ErrorResponse "internal server error"
// @Router /api/v1/image/image_id/{id} [get]

func (h *Handler) GetImageByID(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		NewErrorResponse(c, http.StatusBadRequest, "inavil request params")
		return
	}

	img, err := h.services.Image.GetByID(id)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Header("Content-Disposition", "attachment; filename=image.png")
	c.Data(http.StatusOK, "application/octet-stream", img.Image)
}
