package main

import "fmt"

func main() {
	n = 5
	result := fact(n)
	fmt.Println("Factorial dari", n, "adalah", result)
}

var factorial uint64 = 1
var i int = 1
var n int

func fact(n int) uint64 {
	if n < 0 {
		fmt.Print("erorr")
	} else {
		for i := 1; i <= n; i++ {
			factorial *= uint64(i)
		}
	}
	return factorial
}
