package handler

import (
	"net/http"
	"school21_project1/types"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateClient godoc
//
//	@Summary		Create Client
//	@Description	Create a new client with the provided information
//	@Tags			client
//	@Accept			json
//	@Produce		json
//	@Param			input	body		types.CreateClient		true		"Client Info"
//	@Success		201		{object}	map[string]interface{}	"id":int	"Successful response with client ID"
//	@Failure		400		{object}	ErrorResponse	"invalid request params"
//	@Failure		500		{object}	ErrorResponse	"internal server error"
//	@Router			/client [post]
func (h *Handler) CreateClient(c *gin.Context) {
	var client types.CreateClient

	if err := c.BindJSON(&client); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid request params")
		return
	}

	id, err := h.services.Client.Create(client)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"id": id,
	})
}

// DeleteClientByID godoc
//
//	@Summary		Delete Client by ID
//	@Description	Delete client by ID
//	@Tags			client
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int						true		"Client ID"
//	@Success		200	{object}	map[string]interface{}	"status":	"ok"	"Successful response"
//	@Failure		400	{object}	ErrorResponse	"invalid request params"
//	@Failure		500	{object}	ErrorResponse	"internal server error"
//	@Router			/client/{id} [delete]
func (h *Handler) DeleteClientByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid request params")
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

// FindClientByData godoc
//
//	@Summary		Find Client by data
//	@Description	Find a client by name and surname
//	@Tags			client
//	@Accept			json
//	@Produce		json
//	@Param			name	query		string					true	"Name"
//	@Param			surname	query		string					true	"Surname"
//	@Success		200		{object}	types.ClientDAO			"Successful response with client"
//	@Failure		400		{object}	ErrorResponse	"invalid request params"
//	@Failure		404		{object}	ErrorResponse	"Not found"
//	@Router			/client/find/ [get]
func (h *Handler) FindClientByData(c *gin.Context) {
	name := c.Query("name")
	surname := c.Query("surname")

	if name == "" || surname == "" {
		NewErrorResponse(c, http.StatusBadRequest, "invalid request params")
		return
	}

	output, err := h.services.Client.Find(name, surname)
	if err != nil {
		NewErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, output)
}

// GetAllClients godoc
//
//	@Summary		Get All Clients
//	@Description	Get all clients with pagination parameters
//	@Tags			client
//	@Accept			json
//	@Produce		json
//	@Param			limit	query		int						false			"Limit"
//	@Param			offset	query		int						false			"Offset"
//	@Success		200		{object}	[]types.ClientDAO		""Successful	response	with	clients"
//	@Failure		400		{object}	ErrorResponse	"invalid request params"
//	@Failure		500		{object}	ErrorResponse	"internal server error"
//	@Router			/client [get]
func (h *Handler) GetAllClients(c *gin.Context) {

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

	output, err := h.services.Client.GetAll(limit, offset)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, output)
}

// UpdateClient godoc
//
//	@Summary		Update Client adress
//	@Description	Update client adress information
//	@Tags			client
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int						true		"Client ID"
//	@Param			input	body		types.AdressDTO			true		"Adress info"
//	@Success		200		{object}	StatusResponse	"status":	"ok"	"Successful response"
//	@Failure		400		{object}	ErrorResponse	"invalid request params"
//	@Failure		500		{object}	ErrorResponse	"internal server error"
//	@Router			/client/{id} [patch]
func (h *Handler) UpdateClient(c *gin.Context) {
	var adress types.AdressDTO

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := c.BindJSON(&adress); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Client.Update(id, adress)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, StatusResponse{
		Status: "ok",
	})
}
