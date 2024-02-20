package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// i := 21
	// persen := "%"
	// boolean := true
	// ya :=  "Я(ya)"
	// fmt.Printf("ini adalah i = %v\n", i)
	// fmt.Printf("ini adalah tipe data i= %T\n", i)
	// fmt.Printf("%v\n", persen)
	// fmt.Printf("%v\n", boolean)
	// fmt.Printf("%v\n", boolean)
	// fmt.Printf("%v\n", ya)
	// fmt.Printf("menampilkan nilai base 10: %d\n", i)
	// fmt.Printf("menampilkan nilai base 8: %o\n", i)

	// challange 1
	// var myNum int32 = 50
	// var myNum2 float32 = 51
	// myNumStr := "50"
	// fmt.Printf("ini adalah myNum %v\n", myNum)
	// fmt.Printf("ini adalah myNum2 %v\n", myNum2)
	// fmt.Printf("ini adalah myNumStr %v\n", myNumStr)

	// Challange 2
	// var x int32 = 5
	// var y int32 = 10
	// var z int32 = x + y
	// fmt.Printf("nilai %d\n", z)

	// Challange 3	
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukan nama: ")
	text, _ := reader.ReadString('\n')
	fmt.Println("Hello", text)

	// Challange 4
	people := []string{"Walt", "Jesse", "Skyler", "Saul"}

    fmt.Println("Panjang slice people:", len(people), people)
    people = append(people, "Hank", "Marie")
    fmt.Println("Panjang slice people setelah ditambahkan Hank dan Marie:", len(people), people)
    fmt.Println("Slice people:", people)
}