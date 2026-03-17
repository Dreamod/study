package verify

import (
	"go/study/configs"
	"go/study/pkg/request"
	"go/study/pkg/response"
	"go/study/pkg/verify"
	"log"
	"net/http"
	"time"
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
		//получает входящие данные
		payload, err := request.HandleBody[SendRequest](&w, r)
		if err != nil {
			return
		}

		data := SendResponse{
			Success: false,
			Data:    "Ошибка при отправке письма",
			Errors:  nil,
		}

		//генерируем хэщ
		hash := verify.MakeHash(payload.Email + time.Now().String())

		if len(hash) == 0 {
			log.Fatal(err)
			data.Errors = append(data.Errors, err)
			response.Json(w, data, http.StatusInternalServerError)
		}

		//отправляем письмо
		err = verify.SendEmail(payload.Email, hash)
		if err != nil {
			log.Fatal(err)
			data.Errors = append(data.Errors, err)
			response.Json(w, data, http.StatusInternalServerError)
		}

		//сохраняем хэш в файл
		err = verify.SaveHash(payload.Email, hash)
		if err != nil {
			log.Fatal(err)
			data.Errors = append(data.Errors, err)
			response.Json(w, data, http.StatusInternalServerError)
		}

		//пишем успешный ответ
		data = SendResponse{
			Success: true,
			Data:    "Письмо со ссылкой успешно отправлено",
			Errors:  nil,
		}
		response.Json(w, data, http.StatusOK)
	}
}

func (handler *VerifyHandler) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//получает входящие данные
		hash := r.PathValue("hash")
		if hash == "" {
			return
		}
		result := "Верификация email провалена!"
		//проверим hash в файле
		if verify.VerifyHash(hash) {
			result = "Верификация email успешно завершена!"
			//удаляем файл с хэшем
			verify.RemoveDbData()
		}
		//пишем ответ
		w.Write([]byte(result))
	}
}
