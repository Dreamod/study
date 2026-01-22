package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

func randomHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	n := rand.Intn(6) + 1
	fmt.Fprint(w, n)
}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/random", randomHandler)
	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
