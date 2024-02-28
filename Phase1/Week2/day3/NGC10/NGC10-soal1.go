package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)


func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error:", r)
		}
	}()

	word1, word2 := getInput()

	if !isValidInput(word1) || !isValidInput(word2) {
		panic("Input tidak valid")
	}

	if areAnagrams(word1, word2) {
		fmt.Printf("%s & %s merupakan anagram\n", word1, word2)
	} else {
		fmt.Printf("%s & %s bukan merupakan anagram\n", word1, word2)
	}
}

func getInput() (string, string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan kata pertama: ")
	word1, _ := reader.ReadString('\n')
	fmt.Print("Masukkan kata kedua: ")
	word2, _ := reader.ReadString('\n')
	return strings.TrimSpace(word1), strings.TrimSpace(word2)
}

func isValidInput(word string) bool {
	if len(word) > 10 {
		return false
	}
	for _, char := range word {
		if char <= ' ' || char > '~' {
			return false
		}
	}
	return true
}

func areAnagrams(word1, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}
	word1 = sortString(word1)
	word2 = sortString(word2)
	return word1 == word2
}

func sortString(s string) string {
	sChars := strings.Split(s, "")
	sort.Strings(sChars)
	return strings.Join(sChars, "")
}
