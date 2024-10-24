package main

import (
	"fmt"
)

// Konstanta waktu maksimum (301 menit untuk soal yang tidak terselesaikan)
const MAX_TIME = 301

// Prosedur hitungSkor untuk menghitung total soal yang diselesaikan dan total skor waktu yang dibutuhkan
func hitungSkor(waktu []int, soal *int, skor *int) {
	*soal = 0
	*skor = 0
	for i := 0; i < 8; i++ {
		if waktu[i] < MAX_TIME {
			*soal++
			*skor += waktu[i]
		}
	}
}

func main() {
	var peserta string
	var n int

	fmt.Print("Masukkan jumlah peserta: ")
	fmt.Scan(&n)

	var pemenang string
	soalPemenang := 0
	skorPemenang := 0

	for i := 0; i < n; i++ {
		// Input nama peserta
		fmt.Print("Masukkan nama peserta: ")
		fmt.Scan(&peserta)

		// Input waktu pengerjaan tiap soal
		waktu := make([]int, 8)
		fmt.Println("Masukkan waktu pengerjaan untuk 8 soal (dalam menit):")
		for j := 0; j < 8; j++ {
			fmt.Scan(&waktu[j])
		}

		// Hitung total soal yang diselesaikan dan total skor (waktu)
		var soal, skor int
		hitungSkor(waktu, &soal, &skor)

		// Tentukan pemenang berdasarkan jumlah soal atau total waktu
		if soal > soalPemenang || (soal == soalPemenang && skor < skorPemenang) {
			pemenang = peserta
			soalPemenang = soal
			skorPemenang = skor
		}
	}

	// Cetak hasil
	fmt.Printf("Pemenang: %s, Soal diselesaikan: %d, Nilai: %d\n", pemenang, soalPemenang, skorPemenang)
}
