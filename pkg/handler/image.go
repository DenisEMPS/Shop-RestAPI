package handler

import (
	"net/http"
	"school21_project1/types"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createImageProduct(c *gin.Context) {
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

func (h *Handler) updateImageByID(c *gin.Context) {
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

func (h *Handler) deleteImageByID(c *gin.Context) {
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

func (h *Handler) getImageByProductID(c *gin.Context) {
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

func (h *Handler) getImageByID(c *gin.Context) {
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
