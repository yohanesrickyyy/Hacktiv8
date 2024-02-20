package main

import "fmt"

func main() {
	day2()
}

func day2() {
	//     score := 98

	//     if score >= 100 || score <= 10 {
	//         fmt.Println("Invalid Score!")
	//     } else if score >= 90 {
	//         fmt.Println("Nilai A")
	//     } else if score >= 80 {
	//         fmt.Println("Nilai B")
	//     } else if score >= 70 {
	//         fmt.Println("Nilai C")
	//     } else if score >= 60 {
	//         fmt.Println("Nilai D")
	//     } else {
	//         fmt.Println("Nilai E")
	//     }
	for i := 1; i <= 15; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("Hello World")
		} else if i%5 == 0 {
			fmt.Println("World")
		} else if i%3 == 0 {
			fmt.Println("Hello")
		} else {
			fmt.Println(i)
		}
	}
	
}
