package verify

import (
	"fmt"
	"go/study/configs"
	"log"
	"net/http"
	"net/smtp"

	"github.com/jordan-wright/email"
)

type VerifyHandler struct {
	*configs.Config
}

type VerifyHandlerDeps struct {
	*configs.Config
}

func NewVerifyHandler(router *http.ServeMux, deps VerifyHandlerDeps) {
	handler := &VerifyHandler{
		Config: deps.Config,
	}
	router.HandleFunc("POST /send", handler.Send())
	router.HandleFunc("GET /verify/{hash}", handler.Verify())
}

func (handler *VerifyHandler) Send() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		e := email.NewEmail()
		e.From = "Dreamod Go App <dreamod1@gmail.com>"
		e.To = []string{"dreamod-wm@yandex.ru"}
		e.Subject = "Подтверждение email"
		e.HTML = []byte("<a href=\"https://dreamod.ru:8081/verify\">Подтвердить</a>")

		err := e.Send(handler.Config.Email.Address+":587",
			smtp.PlainAuth("",
				handler.Config.Email.Email,
				handler.Config.Email.Password,
				handler.Config.Email.Address))

		if err != nil {
			log.Fatal(err)
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "Письмо со ссылкой отправлено")
	}
}

func (handler *VerifyHandler) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Verify")
	}
}
