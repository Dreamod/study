package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//создаем каналы
	randomNumsChan := make(chan int, 10)
	resultNumsChan := make(chan int, 10)

	//запускаем горутины
	go makeRandomNumSlice(randomNumsChan)
	go squaringNum(randomNumsChan, resultNumsChan)

	//читаем результирующий канал
	for num := range resultNumsChan {
		fmt.Printf("%d ", num)
	}
}

// создает 10 псевдослучайных целых чисел и отправляет в канал randomNumsChan
func makeRandomNumSlice(randomNumsChan chan int) {
	// создаем пустой слайс на 10 элементов
	randomNumSlice := make([]int, 0, 10)

	//заполняем слайс рандомными числами от 0 до 100
	for i := 0; i < 10; i++ {
		randomNumSlice = append(randomNumSlice, rand.Intn(101))
	}

	//бежим по слайсу и по одному отправляем числа в канал randomNumsChan
	for _, num := range randomNumSlice {
		randomNumsChan <- num
	}

	//канал заполнен - закрываем его
	close(randomNumsChan)

	//wg не нужен т.к. мы точно знаем размерность канала
	return
}

// получает данные из канала randomNumsChan,
// возводит числа в квадрат и записывает в канал resultNumsChan
func squaringNum(randomNumsChan chan int, resultNumsChan chan int) {
	//читаем канал с рандомными числами,
	//возводим в квадрат,
	//пишем в результирующий канал
	for num := range randomNumsChan {
		resultNumsChan <- num * num
	}
	//канал заполнен - закрываем его
	close(resultNumsChan)

	//wg не нужен т.к. мы точно знаем размерность канала
	return
}
