package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}

	wg.Add(3)
	go func() {
		defer wg.Done()
		s := sum(2, 6)
		fmt.Printf("sum:%f\n", s)
	}()

	go func() {
		defer wg.Done()
		f := fact(10)
		fmt.Printf("factorial:%d\n", f)
	}()

	go func() {
		defer wg.Done()
		s := sumSeries(2, 6)
		fmt.Printf("sum of number series:%d\n", s)
	}()

	wg.Wait()
	println("the end")
}

func sum(a, b float64) float64 {
	return a + b
}

func fact(n int) int {
	f := 1
	for i := 1; i <= n; i++ {
		f *= i
	}
	time.Sleep(time.Second * 3)

	return f
}

func sumSeries(a, b int) int {
	sum := 0

	for i := a; i < b; i++ {
		sum += i
	}
	time.Sleep(time.Second * 3)

	return sum
}

func random(n int) int {
	time.Sleep(time.Second * 3)
	return rand.Intn(n)
}
