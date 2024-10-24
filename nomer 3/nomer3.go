package main

import (
	"fmt"
)

// Prosedur cetakDeret untuk mencetak setiap suku dari deret
func cetakDeret(n int) {
	for n != 1 {
		fmt.Printf("%d ", n)
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
	}
	// Cetak 1 sebagai suku terakhir
	fmt.Printf("%d\n", n)
}

func main() {
	// Input
	var n int
	fmt.Print("Masukkan bilangan bulat positif (kurang dari 1000000): ")
	fmt.Scan(&n)

	// Panggil prosedur untuk mencetak deret
	cetakDeret(n)
}
