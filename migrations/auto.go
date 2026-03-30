package main

import (
	"go/study/configs"
	"go/study/internal/link"
	"go/study/pkg/db"
)

func main() {
	conf := configs.LoadConfig()
	db := db.NewDb(conf)
	db.AutoMigrate(&link.Link{})
}
