package handler

import (
	"net/http"
	"school21_project1/types"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateSupplier godoc
// @Summary Create Supplier
// @Description Create a new supplier with the provided information
// @Tags supplier
// @Accept  json
// @Produce  json
// @Param input body types.CreateSupplier true "Supplier Info"
// @Success 201 {object} map[string]interface{} "id":int "Successful response with supplier ID"
// @Failure 400 {object} response.ErrorResponse "invalid request params"
// @Failure 500 {object} response.ErrorResponse "internal server error"
// @Router /api/v1/supplier [post]

func (h *Handler) CreateSupplier(c *gin.Context) {
	var supplier types.CreateSupplier

	err := c.BindJSON(&supplier)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid request params")
		return
	}

	id, err := h.services.Supplier.Create(supplier)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"id": id,
	})
}

// UpdateSupplier godoc
// @Summary Update Supplier Adress
// @Description Update supplier adress information
// @Tags supplier
// @Accept  json
// @Produce  json
// @Param id path int true "Supplier ID"
// @Param input body types.AdressDTO true "Adress info"
// @Success 200 {object} response.StatusResponse "status": "ok" "Successful response"
// @Failure 400 {object} response.ErrorResponse "invalid request params"
// @Failure 500 {object} response.ErrorResponse "internal server error"
// @Router /api/v1/supplier/{id} [patch]

func (h *Handler) UpdateSupplier(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid request params")
		return
	}

	var adress types.AdressDTO
	err = c.BindJSON(&adress)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid request params")
		return
	}

	err = h.services.Supplier.Update(id, adress)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, StatusResponse{
		Status: "ok",
	})
}

// DeleteSupplierByID godoc
// @Summary Delete Supplier by ID
// @Description Delete supplier by supplier ID
// @Tags supplier
// @Accept  json
// @Produce  json
// @Param id path int true "Supplier ID"
// @Success 200 {object} response.StatusResponse "status": "ok" "Successful response"
// @Failure 400 {object} response.ErrorResponse "invalid request params"
// @Failure 500 {object} response.ErrorResponse "internal server error"
// @Router /api/v1/supplier/{id} [delete]

func (h *Handler) DeleteSupplierByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid request params")
		return
	}

	err = h.services.Supplier.Delete(id)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, StatusResponse{
		Status: "ok",
	})
}

// GetAllSuppliers godoc
// @Summary Get All Suppliers
// @Description Get all suppliers with follow up information
// @Tags supplier
// @Accept  json
// @Produce  json
// @Success 200 {object} response.DataResponse "data": []types.SupplierDAO "Successful response with suppliers"
// @Failure 500 {object} response.ErrorResponse "internal server error"
// @Router /api/v1/supplier [get]

func (h *Handler) GetAllSuppliers(c *gin.Context) {
	supplier, err := h.services.Supplier.GetAll()
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, DataResponse{
		Data: supplier,
	})
}

// GetSupplierByID godoc
// @Summary Get Supplier by ID
// @Description Get supplier by ID with follow up information
// @Tags supplier
// @Accept  json
// @Produce  json
// @Param id path int true "Supplier ID"
// @Success 200 {object} types.SupplierDAO "Successful response with supplier"
// @Failure 400 {object} response.ErrorResponse "invalid request params"
// @Failure 404 {object} response.ErrorResponse "Not found"
// @Router /api/v1/supplier/{id} [get]

func (h *Handler) GetSupplierByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid request params")
		return
	}

	supplier, err := h.services.Supplier.GetByID(id)
	if err != nil {
		NewErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, supplier)
}
