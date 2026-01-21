package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// **Описание**: Создайте обработчик для маршрута /form, который читает данные из тела POST-запроса и возвращает размер полученных данных в байтах
//
// **Входные данные**: HTTP POST-запросы с различными данными в теле запроса к маршруту /form
//
// **Выходные данные**: Текстовый ответ в формате "Body size: N bytes", где N - размер тела запроса в байтах
//
// **Ограничения**:
// - Использовать стандартную библиотеку net/http
// - Создать отдельную функцию-обработчик для маршрута /form
// - Зарегистрировать обработчик в кастомном ServeMux
// - Сервер должен слушать порт 8080
// - Использовать io.ReadAll для чтения тела запроса
// - Обработать только POST-запросы
//
// **Примеры**:
// Запрос: POST http://localhost:8080/form с телом "name=John&age=25"
// Ответ: "Body size: 17 bytes"
//
// Запрос: POST http://localhost:8080/form с телом "{\"message\":\"hello\"}"
// Ответ: "Body size: 19 bytes"

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(405)
		return
	}
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	fmt.Fprintln(w, "Body size:", len(body), "bytes")
}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/form", formHandler)
	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
