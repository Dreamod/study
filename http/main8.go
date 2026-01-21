package main

import (
	"fmt"
	"log"
	"net/http"
)

// **Описание**: Создайте обработчик для маршрута /headers, который читает все заголовки запроса и возвращает их количество в текстовом формате
//
// **Входные данные**: HTTP-запросы с различными заголовками к маршруту /headers
//
// **Выходные данные**: Текстовый ответ в формате "Headers count: N", где N - количество заголовков в запросе
//
// **Ограничения**:
// - Использовать стандартную библиотеку net/http
// - Создать отдельную функцию-обработчик для маршрута /headers
// - Зарегистрировать обработчик в кастомном ServeMux
// - Сервер должен слушать порт 8080
// - Подсчитать все заголовки из r.Header
//
// **Примеры**:
// Запрос: GET http://localhost:8080/headers с заголовками User-Agent, Accept, Connection
// Ответ: "Headers count: 3"
//
// Запрос: POST http://localhost:8080/headers с заголовками Content-Type, Authorization
// Ответ: "Headers count: 2"

func headersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Headers count:", len(r.Header))
}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/headers", headersHandler)
	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
