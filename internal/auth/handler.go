package auth

import (
	"fmt"
	"go/study/configs"
	"go/study/pkg/request"
	"go/study/pkg/response"
	"net/http"
)

type AuthHandler struct {
	*configs.Config
}

type AuthHandlerDeps struct {
	*configs.Config
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config: deps.Config,
	}
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		payload, err := request.HandleBody[LoginRequest](&w, r)
		if err != nil {
			return
		}
		data := LoginResponse{
			Token: "1234567890",
		}
		fmt.Println(payload)
		response.Json(w, data, http.StatusOK)
	}
}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//получает входящие данные
		payload, err := request.HandleBody[RegisterRequest](&w, r)
		if err != nil {
			return
		}
		data := RegisterResponse{
			Token: "dujfhgkjdfhg",
			Name:  payload.Name,
		}
		response.Json(w, data, http.StatusOK)
	}
}
