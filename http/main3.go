package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// **Описание**: Создайте кастомный ServeMux, зарегистрируйте в нем обработчик для маршрута /info, который возвращает JSON-ответ с информацией о методе запроса и пути
//
// **Входные данные**: HTTP-запросы к серверу на порту 8080
//
// **Выходные данные**: JSON-ответ в формате {"method": "GET", "path": "/info"} с правильным Content-Type заголовком
//
// **Ограничения**:
// - Использовать http.NewServeMux() для создания кастомного мультиплексора
// - Сервер должен слушать порт 8080
// - Обработчик должен устанавливать Content-Type: application/json
// - Использовать encoding/json для формирования ответа
//
// **Примеры**:
// Запрос: GET http://localhost:8080/info
// Ответ: {"method": "GET", "path": "/info"} с Content-Type: application/json
//
// Запрос: POST http://localhost:8080/info
// Ответ: {"method": "POST", "path": "/info"} с Content-Type: application/json

type Response struct {
	Method string `json:"method"`
	Path   string `json:"path"`
}

func infoHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Method: r.Method,
		Path:   r.URL.Path,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {

	router := http.NewServeMux()
	router.HandleFunc("/info", infoHandler)
	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
