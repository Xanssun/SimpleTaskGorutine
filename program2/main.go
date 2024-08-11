package main

import (
	"fmt"
)

func fenobachi(n int, ch chan int) {
	// алгоритм финобачи
	a, b := 0, 1
	for i := 0; i < n; i++ {
		ch <- a
		a, b = b, a+b
	}
	close(ch)
}

func main() {
	ch := make(chan int)
	go fenobachi(10, ch)

	for num := range ch {
		fmt.Println(num)
	}

}
