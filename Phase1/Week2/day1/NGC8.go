package main

import "fmt"

func main() {
	for i := 1; i < 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Printf("%d-FizzBuzz\n")
			continue
		}
		if i%3 == 0 {
			fmt.Printf("%d-Fizz\n")
			continue
		}
		if i%5 == 0 {
			fmt.Printf("%d-Buzz\n")
			continue
		}
		fmt.Println(i)
	}
}
