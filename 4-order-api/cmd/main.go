package main

import (
	"go/study/4-order-api/configs"
	"go/study/4-order-api/pkg/db"
)

func main() {
	conf := configs.LoadConfig()
	_ = db.NewDb(conf)
}
