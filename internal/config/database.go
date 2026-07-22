package config

import (
	"golang-restapi-big-structure/internal/modules/product"
	"golang-restapi-big-structure/internal/modules/user"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {

	dsn := "root:@(127.0.0.1:3306)/golang_restapi_big_structure?charset=utf8mb4&loc=Local&parseTime=true"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(&user.User{}, &product.Product{}); err != nil {
		panic(err)
	}

	return db, nil
}
