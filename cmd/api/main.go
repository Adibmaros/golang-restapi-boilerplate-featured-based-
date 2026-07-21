package main

import (
	"log"

	"golang-restapi-big-structure/internal/config"
	"golang-restapi-big-structure/internal/modules/product"
	"golang-restapi-big-structure/internal/modules/user"
	"golang-restapi-big-structure/internal/pkg/jwt"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatalf("Gagal koneksi ke database: %v", err)
	}

	// Inisialisasi JWT Service
	jwtSecret := []byte("SECRET_KEY_SUPER_SECRET")
	jwtService := jwt.NewJWTService(jwtSecret)

	// Inisialisasi Layer User (Repository -> Service -> Controller)
	userRepo := user.NewRepository(db)
	userService := user.NewService(userRepo)
	userController := user.NewController(userService)

	// Inisialisasi Layer Product (Repository -> Service -> Controller)
	productRepo := product.NewRepository(db)
	productService := product.NewService(productRepo)
	productController := product.NewController(productService)

	// Setup Router Gin
	r := gin.Default()
	apiGroup := r.Group("/api/v1")

	// Registrasi Routes
	user.UserRoutes(apiGroup, userController, jwtService)
	product.ProductRoutes(apiGroup, productController, jwtService)

	log.Println("Server berjalan di port :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Gagal menjalankan server: %v", err)
	}
}
