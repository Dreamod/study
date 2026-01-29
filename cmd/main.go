package main

import (
	"fmt"
	"go/study/configs"
	"go/study/internal/auth"
	"go/study/internal/verify"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	router := http.NewServeMux()
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: conf,
	})
	verify.NewVerifyHandler(router, verify.VerifyHandlerDeps{
		Config: conf,
	})

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Server is running on port 8081")
	server.ListenAndServe()
}
