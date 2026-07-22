// @title           coba-coba API
// @version         1.0
// @description     API documentation for coba-coba backend service
// @termsOfService  http://swagger.io/terms/

// @contact.name   Adib Maros
// @contact.email  your-email@example.com

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

package main

import (
	"log"
	"os"

	"golang-restapi-big-structure/internal/config"
	"golang-restapi-big-structure/internal/middleware"
	"golang-restapi-big-structure/internal/modules/product"
	"golang-restapi-big-structure/internal/modules/user"
	"golang-restapi-big-structure/internal/pkg/jwt"

	_ "golang-restapi-big-structure/docs" // generated docs package, wajib di-import

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatalf("Gagal koneksi ke database: %v", err)
	}

	// Inisialisasi JWT Service
	jwtSecretStr := os.Getenv("JWT_SECRET")
	if jwtSecretStr == "" {
		jwtSecretStr = "SECRET_KEY_SUPER_SECRET"
	}
	jwtSecret := []byte(jwtSecretStr)
	jwtService := jwt.NewJWTService(jwtSecret)

	// Inisialisasi Layer User (Repository -> Service -> Controller)
	userRepo := user.NewRepository(db)
	userService := user.NewService(userRepo, jwtService)
	userController := user.NewController(userService)

	// Inisialisasi Layer Product (Repository -> Service -> Controller)
	productRepo := product.NewRepository(db)
	productService := product.NewService(productRepo)
	productController := product.NewController(productService)

	// Setup Router Gin
	r := gin.Default()

	// Pasang Middleware CORS
	r.Use(middleware.CORSMiddleware())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	apiGroup := r.Group("/api/v1")

	// Registrasi Routes
	user.UserRoutes(apiGroup, userController, jwtService)
	product.ProductRoutes(apiGroup, productController, jwtService)

	log.Println("Server berjalan di port :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Gagal menjalankan server: %v", err)
	}
}
