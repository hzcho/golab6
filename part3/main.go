package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func generateRandomNumbers(numCh chan<- int, count int) {
	for i := 0; i < count; i++ {
		num := rand.Intn(100)
		numCh <- num
		time.Sleep(time.Second)
	}
}

func checkEvenOdd(resultCh chan<- string, numCh <-chan int) {
	for num := range numCh {
		if num%2 == 0 {
			resultCh <- fmt.Sprintf("%d - чётное", num)
		} else {
			resultCh <- fmt.Sprintf("%d - нечётное", num)
		}
	}
}

func main() {
	numCh := make(chan int)
	resultCh := make(chan string)
	wg := sync.WaitGroup{}

	wg.Add(2)

	go func() {
		defer wg.Done()
		generateRandomNumbers(numCh, 10)
		close(numCh)
	}()
	go func() {
		defer wg.Done()
		checkEvenOdd(resultCh, numCh)
	}()

	go func() {
		for {
			select {
			case result, ok := <-resultCh:
				if ok {
					fmt.Println(result)
				} else {
					close(resultCh)
					fmt.Println("Канал resultCh закрыт.")
					return
				}
			}
		}
	}()

	wg.Wait()
}
