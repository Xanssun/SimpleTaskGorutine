package main

import (
	"fmt"
	"sync"
	"time"
)

type Ttype struct {
	id         int
	cT         string
	fT         string
	taskRESULT []byte
}

func taskCreturer(taskChan chan Ttype, wg *sync.WaitGroup) {
	defer wg.Done()

	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	timeout := time.After(10 * time.Second)

	for {
		select {
		case <-timeout:
			close(taskChan)
			return
		case t := <-ticker.C:
			ft := t.Format(time.RFC3339)
			if t.Nanosecond()%2 > 0 {
				ft = "Some error occurred"
			}
			task := Ttype{cT: ft, id: int(time.Now().Unix())}
			taskChan <- task
		}
	}
}

func taskWorker(taskChan chan Ttype, doneChan chan Ttype, errorChan chan error, wg *sync.WaitGroup) {
	defer wg.Done()

	for task := range taskChan {
		tt, err := time.Parse(time.RFC3339, task.cT)
		if err != nil {
			task.taskRESULT = []byte("failed: invalid creation time format")
		} else {
			if tt.After(time.Now().Add(-20 * time.Second)) {
				task.taskRESULT = []byte("task has been successed")
			} else {
				task.taskRESULT = []byte("failed: task too old")
			}
		}
		task.fT = time.Now().Format(time.RFC3339Nano)

		if string(task.taskRESULT) == "task has been successed" {
			doneChan <- task
		} else {
			errorChan <- fmt.Errorf("Task id %d, Creation time: %s, Result: %s", task.id, task.cT, task.taskRESULT)
		}
	}
}

func printResults(doneTasks *[]Ttype, undoneTasks *[]error, mu *sync.Mutex) {
	mu.Lock()
	defer mu.Unlock()

	fmt.Println("Errors:")
	for _, err := range *undoneTasks {
		fmt.Println(err)
	}

	fmt.Println("Done tasks:")
	for _, task := range *doneTasks {
		fmt.Printf("Task id: %d, Creation time: %s, Finish time: %s, Result: %s\n",
			task.id, task.cT, task.fT, task.taskRESULT)
	}
}

func main() {
	taskChan := make(chan Ttype, 10)
	doneChan := make(chan Ttype, 10)
	errorChan := make(chan error, 10)

	var wg sync.WaitGroup
	var mu sync.Mutex

	doneTasks := []Ttype{}
	undoneTasks := []error{}

	wg.Add(1)
	go taskCreturer(taskChan, &wg)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go taskWorker(taskChan, doneChan, errorChan, &wg)
	}

	go func() {
		for task := range doneChan {
			mu.Lock()
			doneTasks = append(doneTasks, task)
			mu.Unlock()
		}
	}()

	go func() {
		for err := range errorChan {
			mu.Lock()
			undoneTasks = append(undoneTasks, err)
			mu.Unlock()
		}
	}()

	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()

	go func() {
		for range ticker.C {
			printResults(&doneTasks, &undoneTasks, &mu)
		}
	}()

	wg.Wait()

	printResults(&doneTasks, &undoneTasks, &mu)

}
