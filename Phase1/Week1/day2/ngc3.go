package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	//Nomor 1 - Looping 2
	slice1 := []float64{1, 5, 7, 8, 10, 24, 33}
	slice2 := []float64{1.1, 5.4, 6.7, 9.2, 11.3, 25.2, 33.1}

	hitungDanTampilkanStatistik(slice1)
	hitungDanTampilkanStatistik(slice2)

	// Nomor 2 - Palindrome
	// kata := "katak"
	// fmt.Println("Apakah kata", kata, "palindrome?", cekPalindrome(kata))

	// kata = "hello"
	// fmt.Println("Apakah kata", kata, "palindrome?", cekPalindrome(kata))

	//Nomor 3 Logic 2 - XOXO
	kata := "xoxo"
	fmt.Println("Apakah jumlah karakter x dan o sama?", cekJumlahXO(kata))

	kata = "xoxoxo"
	fmt.Println("Apakah jumlah karakter x dan o sama?", cekJumlahXO(kata))

	kata = "xoxoxoox"
	fmt.Println("Apakah jumlah karakter x dan o sama?", cekJumlahXO(kata))

	// Nomor 3 Logic 3 - XOXO
	// Membuat slice integer
	slice := []int{9, 4, 7, 2, 1, 6}

	// Mengurutkan slice dari besar ke kecil
	for i := 0; i < len(slice); i++ {
		for j := i + 1; j < len(slice); j++ {
			if slice[i] < slice[j] {
				// Menukar posisi elemen jika elemen pertama lebih kecil dari elemen kedua
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}

	// Menampilkan hasil setelah diurutkan
	fmt.Println("Slice setelah diurutkan:", slice)

	// Nomor 4 Logic 4 - Asterisk Level 1
	rows1 := 5

	fmt.Println("Asterik")
	for i := 0; i < rows1; i++ {
		fmt.Println("*")
	}
	fmt.Println()

	// Nomor 4 Logic 5 - Asterisk Level 2
	rows2 := 5

	fmt.Println("Asterik2")
	for i := 0; i < rows2; i++ {
		for j := 0; j < rows2; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	fmt.Println()

	// NO 4 Logic 6 - Asterisk Level 3
	rows3 := 5

	fmt.Println("Asterik 3")
	for i := 0; i < rows3; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	fmt.Println()

	// NO 4 : Logic 7 - Asterisk Level 4
	rows4 := 5

	fmt.Println("ASterik4")
	for i := 0; i < rows4; i++ {
		for j := rows4 - i; j > 0; j-- {
			fmt.Print("*")
		}
		fmt.Println()
	}
	fmt.Println()

}

// ! Logic
// Nomor 1 -- Logic Looping
func hitungDanTampilkanStatistik(slice []float64) {
	var total, median float64

	// Menghitung jumlah
	for _, nilai := range slice {
		total += nilai
	}

	// Menghitung rata-rata
	rataRata := total / float64(len(slice))

	// Menghitung median
	sort.Float64s(slice)
	panjang := len(slice)
	if panjang%2 == 0 {
		median = (slice[panjang/2-1] + slice[panjang/2]) / 2
	} else {
		median = slice[panjang/2]
	}

	// Menampilkan hasil
	fmt.Printf("Slice: %v\n", slice)
	fmt.Printf("Rata-rata: %.2f\n", rataRata)
	fmt.Printf("Jumlah: %.2f\n", total)
	fmt.Printf("Median: %.2f\n", median)
	fmt.Println()
}

// Nomor 2 - Logic Palindrome
func cekPalindrome(kata string) bool {
	// Mengonversi kata menjadi huruf kecil dan menghapus spasi
	kata = strings.ToLower(strings.ReplaceAll(kata, " ", " "))

	i := 0
	j := len(kata) - 1

	// Memeriksa apakah kata adalah palindrome
	for i < j {
		if kata[i] != kata[j] {
			return false
		}
		i++
		j--
	}
	return true
}

// Nomor 3 - Logic XOXO
func cekJumlahXO(kata string) bool {
	// Mengonversi kata menjadi huruf kecil untuk memudahkan perbandingan
	kata = strings.ToLower(kata)

	// Menghitung jumlah karakter 'x' dan 'o'
	jumlahX := strings.Count(kata, "x")
	jumlahO := strings.Count(kata, "o")

	// Memeriksa apakah jumlah karakter 'x' dan 'o' sama
	return jumlahX == jumlahO
}
