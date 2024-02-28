package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error:", r)
		}
	}()

	operator := getOperator()
	num1 := getNumber("masukan angka:")
	num2 := getNumber("masukan angka:")

	switch operator {
	case "penjumlahan":
		result := num1 + num2
		fmt.Printf("Hasil penjumlahan %d dan %d adalah %d\n", num1, num2, result)
	case "pengurangan":
		result := num1 - num2
		fmt.Printf("Hasil pengurangan %d dan %d adalah %d\n", num1, num2, result)
	case "perkalian":
		result := num1 * num2
		fmt.Printf("Hasil perkalian %d dan %d adalah %d\n", num1, num2, result)
	case "pembagian":
		if num2 == 0 {
			panic("Pembagian oleh nol tidak diperbolehkan")
		}
		result := float64(num1) / float64(num2)
		fmt.Printf("Hasil pembagian %d dan %d adalah %.2f\n", num1, num2, result)
	default:
		panic("Operasi aritmatika tidak valid")
	}
}

func getOperator() string {
	fmt.Println("Pilih operasi aritmatika:")
	fmt.Println("a) penjumlahan (+)")
	fmt.Println("b) pengurangan (-)")
	fmt.Println("c) perkalian (*)")
	fmt.Println("d) pembagian (/)")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("> ")
	operator, _ := reader.ReadString('\n')
	operator = strings.TrimSpace(operator)

	switch operator {
	case "a":
		return "penjumlahan"
	case "b":
		return "pengurangan"
	case "c":
		return "perkalian"
	case "d":
		return "pembagian"
	default:
		panic("Operasi aritmatika tidak valid")
	}
}

func getNumber(prompt string) int {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		num, err := strconv.Atoi(input)
		if err == nil {
			return num
		}
		fmt.Println("Input tidak valid, masukkan angka!")
	}
}
