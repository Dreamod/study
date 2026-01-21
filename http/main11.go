package main

import (
	"fmt"
	"log"
	"net/http"
)

// **Описание**: Создайте обработчик для маршрута /method, который проверяет HTTP-метод запроса и возвращает разные сообщения в зависимости от метода
//
// **Входные данные**: HTTP-запросы с различными методами (GET, POST, PUT, DELETE) к маршруту /method
//
// **Выходные данные**: Текстовый ответ в зависимости от метода:
// - GET: "This is a GET request"
// - POST: "This is a POST request"
// - PUT: "This is a PUT request"
// - DELETE: "This is a DELETE request"
// - Другие методы: "Unsupported method"
//
// **Ограничения**:
// - Использовать стандартную библиотеку net/http
// - Создать отдельную функцию-обработчик для маршрута /method
// - Зарегистрировать обработчик в кастомном ServeMux
// - Сервер должен слушать порт 8080
// - Использовать r.Method для определения HTTP-метода
//
// **Примеры**:
// Запрос: GET http://localhost:8080/method
// Ответ: "This is a GET request"
//
// Запрос: POST http://localhost:8080/method
// Ответ: "This is a POST request"
//
// Запрос: PATCH http://localhost:8080/method
// Ответ: "Unsupported method"

func methodHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" || r.Method == "POST" || r.Method == "PUT" || r.Method == "DELETE" {
		fmt.Fprintln(w, "This is a "+r.Method+" request")
	} else {
		fmt.Fprintln(w, "Unsupported method")
	}
}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/method", methodHandler)
	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
