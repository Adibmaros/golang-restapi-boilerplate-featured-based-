package user

import (
	"golang-restapi-big-structure/internal/middleware"
	"golang-restapi-big-structure/internal/pkg/jwt"

	"github.com/gin-gonic/gin"
)

func UserRoutes(rg *gin.RouterGroup, controller *controller, jwtService jwt.Service) {
	users := rg.Group("/users")

	users.POST("/register", controller.RegisterUser)

	protected := users.Group("/")
	protected.Use(middleware.AuthMiddleware(jwtService))

	protected.GET("/", controller.GetAllUsers)
	protected.GET("/profile", controller.GetUserByID)
}
