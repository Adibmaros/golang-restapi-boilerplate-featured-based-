package product

import (
	"golang-restapi-big-structure/internal/middleware"
	"golang-restapi-big-structure/internal/pkg/jwt"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(rg *gin.RouterGroup, controller *controller, jwtService jwt.Service) {
	products := rg.Group("/products")

	// Public routes
	products.GET("/", controller.GetAllProducts)
	products.GET("/:id", controller.GetProductByID)

	// Protected routes
	protected := products.Group("/")
	protected.Use(middleware.AuthMiddleware(jwtService))
	protected.POST("/", controller.CreateProduct)
}
