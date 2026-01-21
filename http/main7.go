package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"
)

// **Описание**: Создайте программу, которая отправляет HTTP GET-запрос к локальному серверу и выводит полученный статус-код и тело ответа
//
// **Входные данные**: URL для запроса: "http://localhost:8080/test"
//
// **Выходные данные**: Статус-код ответа и содержимое тела ответа в консоль
//
// **Ограничения**:
// - Использовать стандартную библиотеку net/http для отправки запроса
// - Обработать возможные ошибки при выполнении запроса
// - Закрыть тело ответа после чтения
// - Вывести статус-код и содержимое тела в консоль
//
// **Примеры**:
// Запрос к http://localhost:8080/test
// Вывод:
// Status: 200
// Body: Hello from server

const Port = "8080"

func main() {
	url := flag.String("url", "http://localhost:"+Port+"/test", "url")
	rCh := make(chan string, 1)
	sCh := make(chan struct{})
	flag.Parse()

	router := http.NewServeMux()
	router.HandleFunc("/test", test)
	server := &http.Server{Handler: router}
	listener, err := net.Listen("tcp", ":"+Port)
	if err != nil {
		log.Fatal(err)
	}

	go startServer(sCh, server, listener)
	<-sCh

	go sendRequest(url, rCh)

	response := <-rCh
	fmt.Println(response)

	if err := server.Shutdown(context.Background()); err != nil {
		// Error from closing listeners, or context timeout:
		log.Printf("HTTP server Shutdown: %v", err)
	}
}

// отправляем запрос
func sendRequest(url *string, rCh chan string) {
	resp, err := http.Get(*url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	rCh <- " Status code:  " + strconv.Itoa(resp.StatusCode)
	close(rCh)
}

// стартуем сервер
func startServer(ready chan struct{}, server *http.Server, listener net.Listener) {
	close(ready)
	server.Serve(listener)
}

// роут status
func test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Server is running on port 8080")
}
