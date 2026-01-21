package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// **Описание**: Создайте обработчик для маршрута /timeout, который устанавливает таймаут на запись ответа и возвращает сообщение о времени ожидания
//
// **Входные данные**: HTTP GET-запросы к маршруту /timeout
//
// **Выходные данные**: Текстовый ответ "Request processed with timeout handling" со статус-кодом 200
//
// **Ограничения**:
// - Использовать стандартную библиотеку net/http
// - Создать отдельную функцию-обработчик для маршрута /timeout
// - Зарегистрировать обработчик в кастомном ServeMux
// - Сервер должен слушать порт 8080
// - Использовать time.Sleep(100 * time.Millisecond) для имитации обработки
// - Добавить логирование начала и завершения обработки запроса
//
// **Примеры**:
// Запрос: GET http://localhost:8080/timeout
// Ответ: "Request processed with timeout handling" со статус-кодом 200
//
// В консоли должны появиться сообщения:
// "Processing request..."
// "Request completed"

func timeoutHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	fmt.Fprintln(w, "Request processed with timeout handling")
	fmt.Println("Processing request...")
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Request completed")
}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/timeout", timeoutHandler)
	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
