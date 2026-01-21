package main

import (
	"fmt"
	"net/http"
	"time"
)

type Handler struct{}

func NewHandler(router *http.ServeMux) {
	handler := &Handler{}
	router.HandleFunc("/getTime", handler.getTime())
}

func (handler *Handler) getTime() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(time.Now())
	}
}
