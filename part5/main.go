package main

import (
	"fmt"
	"sync"
)

type Request struct {
	Operation string
	A         float64
	B         float64
}

func calculator(requests <-chan Request, resChan chan<- float64, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(resChan)
	for req := range requests {
		var result float64
		switch req.Operation {
		case "+":
			result = req.A + req.B
		case "-":
			result = req.A - req.B
		case "*":
			result = req.A * req.B
		case "/":
			if req.B != 0 {
				result = req.A / req.B
			} else {
				fmt.Println("Ошибка: деление на ноль")
				result = 0
			}
		default:
			fmt.Println("Ошибка: неизвестная операция")
			result = 0
		}
		resChan <- result
	}
}

func main() {
	requests := make(chan Request)
	resultChan := make(chan float64)
	var wg sync.WaitGroup

	wg.Add(1)
	go calculator(requests, resultChan, &wg)

	go func() {
		requests <- Request{Operation: "+", A: 1, B: 1}
		requests <- Request{Operation: "-", A: 1, B: 1}
		requests <- Request{Operation: "*", A: 2, B: 3}
		requests <- Request{Operation: "/", A: 5, B: 2}
		requests <- Request{Operation: "/", A: 5, B: 0}
		close(requests)
	}()

	for v := range resultChan {
		fmt.Println(v)
	}

	wg.Wait()
}
