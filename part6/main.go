package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

type Task struct {
	Line   string
	Result chan string
}

func worker(tasks <-chan Task, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {
		reversed := reverseString(task.Line)
		task.Result <- reversed
	}
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	var numWorkers int
	fmt.Print("Введите количество воркеров: ")
	fmt.Scan(&numWorkers)

	tasks := make(chan Task, numWorkers)
	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(tasks, &wg)
	}

	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Ошибка при получении текущего каталога:", err)
		return
	}

	file, err := os.Open(currentDir + "/part6/input.txt")
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	results := make([]string, 0)

	for scanner.Scan() {
		resultChan := make(chan string)
		tasks <- Task{Line: scanner.Text(), Result: resultChan}
		results = append(results, <-resultChan)
	}

	close(tasks)

	wg.Wait()

	outputFile, err := os.Create(currentDir + "/part6/output.txt")
	if err != nil {
		fmt.Println("Ошибка при создании файла:", err)
		return
	}
	defer outputFile.Close()

	for _, result := range results {
		_, err := outputFile.WriteString(result + "\n")
		if err != nil {
			fmt.Println("Ошибка при записи в файл:", err)
			return
		}
	}

	fmt.Println("Обработка завершена. Результаты записаны в output.txt.")
}
