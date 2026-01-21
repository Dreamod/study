package main

import (
	"fmt"
	"log"
	"net/http"
)

// **Описание**: Создайте обработчик для маршрута /query, который читает параметры запроса из URL и возвращает их количество в текстовом формате
//
// **Входные данные**: HTTP-запросы с различными query параметрами к маршруту /query
//
// **Выходные данные**: Текстовый ответ в формате "Query parameters count: N", где N - количество параметров в запросе
//
// **Ограничения**:
// - Использовать стандартную библиотеку net/http
// - Создать отдельную функцию-обработчик для маршрута /query
// - Зарегистрировать обработчик в кастомном ServeMux
// - Сервер должен слушать порт 8080
// - Подсчитать все параметры из r.URL.Query()
//
// **Примеры**:
// Запрос: GET http://localhost:8080/query?name=John&age=25&city=Moscow
// Ответ: "Query parameters count: 3"
//
// Запрос: GET http://localhost:8080/query?search=golang
// Ответ: "Query parameters count: 1"

func queryHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Query parameters count:", len(r.URL.Query()))
}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/query", queryHandler)
	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
