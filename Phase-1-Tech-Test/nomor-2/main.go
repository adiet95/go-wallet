package main

import "fmt"

func calculateChange(totalBelanja, uangDibayar int) interface{} {
	if uangDibayar < totalBelanja {
		return false
	}
	kembalian := uangDibayar - totalBelanja
	kembalian = (kembalian / 100) * 100

	pecahan := []int{100000, 50000, 20000, 10000, 5000, 2000, 1000, 500, 200, 100}

	fmt.Printf("Kembalian yang dibulatkan: Rp %d\n", kembalian)
	fmt.Println("Pecahan yang diberikan:")

	for _, p := range pecahan {
		jumlahPecahan := kembalian / p
		if jumlahPecahan > 0 {
			fmt.Printf("%d lembar : %d \n", jumlahPecahan, p)
			kembalian -= jumlahPecahan * p
		}
	}
	return true
}

func main() {
	var totalBelanja, uangDibayar int
	fmt.Print("Total belanja seorang customer: Rp ")
	fmt.Scan(&totalBelanja)
	fmt.Print("Pembeli membayar: Rp ")
	fmt.Scan(&uangDibayar)

	result := calculateChange(totalBelanja, uangDibayar)
	if result == false {
		fmt.Println("Uang yang dibayarkan kurang dari total belanja!")
	}
}
