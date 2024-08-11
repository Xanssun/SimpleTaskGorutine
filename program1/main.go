package main

import (
	"fmt"
	"sync"
)

var count int
var wg sync.WaitGroup
var mutex sync.Mutex

func increment() {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		mutex.Lock()
		count++
		mutex.Unlock()
	}

}

func main() {

	wg.Add(5)

	for i := 0; i < 5; i++ {
		go increment()
	}

	wg.Wait()

	fmt.Println("Count:", count)

}
