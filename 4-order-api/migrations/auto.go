package main

import (
	"go/study/4-order-api/configs"
	"go/study/4-order-api/internal/product"
	"go/study/4-order-api/pkg/db"
)

func main() {
	conf := configs.LoadConfig()
	db := db.NewDb(conf)
	db.AutoMigrate(&product.Product{})
}
