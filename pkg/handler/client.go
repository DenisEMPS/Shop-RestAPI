package handler

import (
	"net/http"
	"school21_project1/types"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createClient(c *gin.Context) {
	var client types.Client
	if err := c.BindJSON(&client); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Client.Create(client)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) deleteClient(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Client.Delete(id); err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, StatusResponse{
		Status: "ok",
	})
}
func (h *Handler) findClientByName(c *gin.Context) {
	name := c.Query("name")
	surname := c.Query("surname")

	if name == "" || surname == "" {
		NewErrorResponse(c, http.StatusBadRequest, "invalid request params")
		return
	}

	output, err := h.services.Client.Find(name, surname)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, output)
}

func (h *Handler) getAllClients(c *gin.Context) {
	limit := c.Query("limit")
	offset := c.Query("offset")

	output, err := h.services.Client.GetAll(limit, offset)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, output)
}

func (h *Handler) updateClientEmail(c *gin.Context) {
	var adress types.Adress

	id := c.Param("id")

	if err := c.BindJSON(&adress); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.services.Client.Update(id, adress)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, StatusResponse{
		Status: "ok",
	})
}
