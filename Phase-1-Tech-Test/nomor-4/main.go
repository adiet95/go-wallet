package main

import (
	"fmt"
	"math"
	"time"
)

// Fungsi untuk memeriksa apakah karyawan dapat mengambil cuti pribadi
func canTakeLeave(joinDate, leaveDate time.Time, cutiBersama int) (bool, string) {
	// Cuti kantor tetap adalah 14 hari per tahun
	totalCutiKantor := 14
	// Menghitung jumlah cuti pribadi setelah dikurangi cuti bersama
	cutiPribadi := totalCutiKantor - cutiBersama

	// Hitung durasi cuti dalam hari
	duration := leaveDate.Sub(joinDate).Hours() / 24

	// 180 hari pertama: Karyawan baru tidak bisa mengambil cuti pribadi
	if duration < 180 {
		return false, "Karyawan baru tidak dapat mengambil cuti pribadi dalam 180 hari pertama."
	}

	// Menghitung sisa hari cuti di tahun pertama
	// Hari yang tersedia setelah 180 hari pertama
	daysAvailable := 365 - int(duration)

	// Menghitung jumlah cuti pribadi yang bisa diambil berdasarkan hari yang tersedia
	// Pembulatan kebawah dengan menggunakan math.Floor
	cutiPribadiMax := int(math.Floor(float64(daysAvailable) / 365 * float64(cutiPribadi)))

	// Jika cuti yang diminta lebih besar dari jumlah yang tersedia
	if leaveDate.Sub(joinDate).Hours()/24 > float64(cutiPribadiMax) {
		return false, fmt.Sprintf("Cuti pribadi yang diminta melebihi kuota yang tersedia: %d hari.", cutiPribadiMax)
	}

	// Memeriksa apakah durasi cuti tidak melebihi 3 hari berturut-turut
	if leaveDate.Sub(joinDate).Hours()/24 > 3 {
		return false, "Durasi cuti pribadi tidak boleh lebih dari 3 hari berturut-turut."
	}

	// Jika semua syarat terpenuhi, karyawan dapat mengambil cuti
	return true, "Cuti pribadi dapat diambil."
}

func main() {
	// Contoh input: Jumlah Cuti Bersama, Tanggal join, Tanggal rencana cuti, Durasi cuti (dalam hari)
	var cutiBersama int
	var joinDateStr, leaveDateStr string

	// Input jumlah cuti bersama dan tanggal
	fmt.Print("Masukkan jumlah cuti bersama: ")
	fmt.Scan(&cutiBersama)
	fmt.Print("Masukkan tanggal join (format: YYYY-MM-DD): ")
	fmt.Scan(&joinDateStr)
	fmt.Print("Masukkan tanggal rencana cuti (format: YYYY-MM-DD): ")
	fmt.Scan(&leaveDateStr)

	// Parsing string tanggal menjadi tipe time.Time
	joinDate, _ := time.Parse("2006-01-02", joinDateStr)
	leaveDate, _ := time.Parse("2006-01-02", leaveDateStr)

	// Memanggil fungsi untuk mengecek apakah karyawan bisa mengambil cuti
	canTake, reason := canTakeLeave(joinDate, leaveDate, cutiBersama)

	// Menampilkan hasil
	if canTake {
		fmt.Println("True: " + reason)
	} else {
		fmt.Println("False: " + reason)
	}
}
