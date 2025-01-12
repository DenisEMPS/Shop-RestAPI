package handler

import (
	"school21_project1/pkg/service"

	_ "school21_project1/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api/v1/")
	{
		client := api.Group("/client")
		{
			client.POST("/", h.CreateClient)
			client.DELETE("/:id", h.DeleteClientByID)
			client.GET("/find/", h.FindClientByData)
			client.GET("/", h.GetAllClients)
			client.PATCH("/:id", h.UpdateClient)
		}

		product := api.Group("/product")
		{
			product.POST("/", h.CreateProduct)
			product.PATCH("/:id", h.UpdateClient)
			product.GET("/", h.GetAllProducts)
			product.GET("/:id", h.GetProductByID)
			product.DELETE("/:id", h.DeleteProductByID)
		}

		supplier := api.Group("/supplier")
		{
			supplier.POST("/", h.CreateSupplier)
			supplier.PATCH("/:id", h.UpdateSupplier)
			supplier.DELETE("/:id", h.DeleteSupplierByID)
			supplier.GET("/", h.GetAllSuppliers)
			supplier.GET("/:id", h.GetSupplierByID)
		}

		image := api.Group("/image")
		{
			image.POST("/", h.CreateImageProduct)
			image.PATCH("/:id", h.UpdateImageByID)
			image.DELETE("/:id", h.DeleteImageByID)
			image.GET("/product_id/:id", h.GetImageByProductID)
			image.GET("/image_id/:id", h.GetImageByID)
		}
	}
	return router
}
