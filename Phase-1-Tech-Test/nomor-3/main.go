package main

import "fmt"

func isValidString(s string) bool {
	// Stack untuk menyimpan karakter pembuka
	stack := []rune{}

	// Menelusuri setiap karakter dalam string
	for _, char := range s {
		// Jika karakter adalah pembuka
		if char == '<' || char == '{' || char == '[' {
			stack = append(stack, char) // Masukkan ke stack
		} else if char == '>' || char == '}' || char == ']' {
			// Jika karakter adalah penutup, periksa apakah ada pembuka yang sesuai
			if len(stack) == 0 {
				// Tidak ada pembuka, langsung invalid
				return false
			}
			top := stack[len(stack)-1] // Ambil pembuka paling atas di stack

			// Periksa apakah penutup sesuai dengan pembuka
			if (char == '>' && top != '<') || (char == '}' && top != '{') || (char == ']' && top != '[') {
				// Penutup tidak sesuai dengan pembuka, invalid
				return false
			}

			// Jika sesuai, keluarkan pembuka dari stack
			stack = stack[:len(stack)-1]
		}
	}

	// Jika stack kosong, artinya semua pembuka sudah terpasang dengan penutupnya
	return len(stack) == 0
}

func main() {
	// Input string
	var input string
	fmt.Print("Masukkan string: ")
	fmt.Scan(&input)

	// Memvalidasi string
	if isValidString(input) {
		fmt.Println("String valid!")
	} else {
		fmt.Println("String tidak valid!")
	}
}
