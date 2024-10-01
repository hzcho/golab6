package main

import (
	"fmt"
	"sync"
)

func fibonacci(n int, ch chan<- int) {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		ch <- a
		a, b = b, a+b
	}
	close(ch)
}

func main() {
	ch := make(chan int)
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		fibonacci(10, ch)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range ch {
			fmt.Println(v)
		}
	}()

	wg.Wait()
	fmt.Println("the end")
}
