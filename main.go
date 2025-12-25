package main

import (
	"fmt"
	"time"
)

func main() {
	go printMessage()
	time.Sleep(1 * time.Second)
}

func printMessage() {
	fmt.Println("Hello from goroutine!")
}
