package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	source := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(source)
	n := rnd.Intn(100) + 1

	inputChannel := make(chan int)
	resultChannel := make(chan int)
	done := make(chan struct{})

	wg.Add(1)
	go factorial(inputChannel, resultChannel, done, &wg)

	sendNumbers(inputChannel, n)
	close(inputChannel)

	go func() {
		wg.Wait()
		close(done)
	}()

	for {
		select {
		case res, ok := <-resultChannel:
			if !ok {
				return
			}
			fmt.Println(res)
		case <-done:
			return
		}
	}
}

func sendNumbers(ch chan<- int, n int) {
	for i := 1; i <= n; i++ {
		ch <- i
	}
}

func factorial(input <-chan int, resultChannel chan<- int, done <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(resultChannel)

	for {
		select {
		case num, ok := <-input:
			if !ok {
				return
			}
			resultChannel <- calcFactorial(num)
		case <-done:
			return
		}
	}
}

func calcFactorial(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}
