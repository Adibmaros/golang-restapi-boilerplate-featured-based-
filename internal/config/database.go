package config

import (
	"fmt"
	"os"

	"golang-restapi-big-structure/internal/modules/product"
	"golang-restapi-big-structure/internal/modules/user"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists && value != "" {
		return value
	}
	return fallback
}

func ConnectDB() (*gorm.DB, error) {
	dbHost := getEnv("DB_HOST", "127.0.0.1")
	dbPort := getEnv("DB_PORT", "3306")
	dbUser := getEnv("DB_USER", "root")
	dbPassword := getEnv("DB_PASSWORD", "")
	dbName := getEnv("DB_NAME", "golang_restapi_big_structure")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&loc=Local&parseTime=true",
		dbUser, dbPassword, dbHost, dbPort, dbName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&user.User{}, &product.Product{}); err != nil {
		return nil, err
	}

	return db, nil
}
