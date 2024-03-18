package main

import (
	"fmt"
	"sync"
)

func printNumbers() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}

func printLetters() {
	for ch := 'a'; ch <= 'j'; ch++ {
		fmt.Printf("%c\n", ch)
	}
}

func main() {
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		printNumbers()
	}()

	go func() {
		defer wg.Done()
		printLetters()
	}()
	wg.Wait()
}
