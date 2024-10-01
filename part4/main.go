package main

import (
	"fmt"
	"sync"
	"time"
)

var counter int
var mu sync.Mutex

func increment(wg *sync.WaitGroup) {
	defer wg.Done()

	mu.Lock()
	time.Sleep(time.Millisecond * 10)
	counter++
	mu.Unlock()
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			increment(&wg)
			fmt.Println(i)
		}()
	}

	wg.Wait()
	fmt.Println("Итоговое значение счетчика:", counter)
}
