package handler

import (
	"school21_project1/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	client := router.Group("/client")
	{
		client.POST("/", h.createClient)
		client.DELETE("/:id", h.deleteClient)
		client.GET("/find/", h.findClientByName)
		client.GET("/", h.getAllClients)
		client.PATCH("/:id", h.updateClientEmail)
	}

	product := router.Group("/product")
	{
		product.POST("/", h.createProduct)
		product.PATCH("/:id", h.updateProduct)
		product.GET("/", h.getAllProducts)
		product.GET("/:id", h.getProductByID)
		product.DELETE("/:id", h.deleteProductByID)
	}

	supplier := router.Group("/supplier")
	{
		supplier.POST("/", h.createSupplier)
		supplier.PATCH("/:id", h.updateSupplier)
		supplier.DELETE("/:id", h.deleteSupplierByID)
		supplier.GET("/", h.getAllSuppliers)
		supplier.GET("/:id", h.getSupplierByID)
	}

	image := router.Group("/image")
	{
		image.POST("/", h.createImageProduct)
		image.PATCH("/:id", h.updateImageByID)
		image.DELETE("/:id", h.deleteImageByID)
		image.GET("/product_id/:id", h.getImageByProductID)
		image.GET("/image_id/:id", h.getImageByID)
	}

	return router
}
