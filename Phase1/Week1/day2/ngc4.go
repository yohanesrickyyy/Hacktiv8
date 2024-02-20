package main

import (
	"fmt"
	"strings"
)

func main() {
	// Soal 1
	fmt.Println(AlayGen("hello", "world", "zzz","asik"))
	// Soal 2
	n := 6
	output := fibo(n)
	fmt.Printf("berdasarkan nomor %d ini, nomor fibonaccinya adalah %d\n", n, output)
}

// ! Soal 1 Logic
func AlayGen(words ...string) string {
	alayMap := map[string]string{
		"a": "4",
		"e": "3",
		"i": "!",
		"l": "1",
		"n": "N",
		"s": "5",
		"x": "*",
	}

	// Menggabungkan kata-kata menjadi satu kalimat
	result := strings.Join(words, " ")

	// merubah huruf normal ke alay
	for key, value := range alayMap {
		result = strings.ReplaceAll(result, key, value)
	}

	return result
}

// ! SOAL 2 Logic
func fibo(n int) int {
	fib := make([]int, n+1)
	fib[0] = 0
	fib[1] = 1

	for i := 2; i <= n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}
	return fib[n]
}
