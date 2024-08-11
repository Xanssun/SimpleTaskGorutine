package main

import "fmt"

func multiplication(ch chan int, data interface{}) {
	// алгоритм умножения
	for _, num := range data.([]int) {
		ch <- num * 2
	}
	close(ch)

}

func main() {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8}
	ch := make(chan int)

	go multiplication(ch, data)

	for num := range ch {
		fmt.Println(num)
	}
}
