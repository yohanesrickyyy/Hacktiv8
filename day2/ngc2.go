package main

import (
	"bufio"
	"fmt"
	"strconv"

	//"math/rand" // soal no 1
	"os"
	//"time" soal no 1
)

func main() {
	/*
		// Menginisialisasi generator angka acak
		rand.Seed(time.Now().UnixNano())

		// Menerima input nama dari pengguna
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Masukkan nama Anda: ")
		nama, _ := reader.ReadString('\n')
		nama = nama[:len(nama)-1]

		// random number
		randomNumber := rand.Intn(100)

		// pesan
		if randomNumber > 80 {
			fmt.Printf("Selamat %s, anda sangat beruntung! Angka anda: %d\n", nama, randomNumber)
		} else if randomNumber <= 80 && randomNumber > 60 {
			fmt.Printf("Selamat %s, anda beruntung! Angka anda: %d\n", nama, randomNumber)
		} else if randomNumber <= 60 && randomNumber > 40 {
			fmt.Printf("Mohon maaf %s, anda kurang beruntung. Angka anda: %d\n", nama, randomNumber)
		} else {
			fmt.Printf("Mohon maaf %s, anda sial. Angka anda: %d\n", nama, randomNumber)
		}
	*/
	// Conditional 2
	// Membaca input nama dari pengguna
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan nama Anda: ")
	nama, _ := reader.ReadString('\n')

	// Membaca input umur dari pengguna
	fmt.Print("Masukkan umur Anda: ")
	umurString, _ := reader.ReadString('\n')

	// Mengonversi string umur menjadi integer
	umur, err := strconv.Atoi(umurString)
	if err != nil {
		fmt.Println("Error: Umur harus berupa angka")
		return
	}

	// Memeriksa dan menampilkan pesan berdasarkan umur
	if umur < 0 || umur > 100 {
		fmt.Println("Error: Umur harus berada dalam rentang 0-100", nama)
		return
	} else if umur > 18 {
		fmt.Println("Silahkan masuk")
	} else {
		fmt.Println("Dilarang masuk, maksimal umur 19")
	}
}


