package main

import (
	"fmt"
	"strings"
)

func findMatchingStrings(N int, stringsList []string) interface{} {
	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			if strings.ToLower(stringsList[i]) == strings.ToLower(stringsList[j]) {
				return fmt.Sprintf("String yang cocok ditemukan: %d %d", i+1, j+1)
			}
		}
	}
	// Jika tidak ada kecocokan
	return false
}

func main() {
	// Input jumlah string
	var N int
	fmt.Print("Masukkan jumlah string: ")
	fmt.Scan(&N)

	// Input string
	stringsList := make([]string, N)
	fmt.Println("Masukkan string:")
	for i := 0; i < N; i++ {
		fmt.Scan(&stringsList[i])
	}

	// Menjalankan fungsi untuk menemukan kecocokan
	result := findMatchingStrings(N, stringsList)
	// Menampilkan hasil
	fmt.Println(result)
}
