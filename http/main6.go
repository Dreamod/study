package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/status", statusHandler)
	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func statusHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Server is running on port 8080")
}
