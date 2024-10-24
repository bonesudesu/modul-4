package main

import (
	"fmt"
)

// Fungsi untuk menghitung faktorial
func faktorial(n int) int {
	if n == 0 {
		return 1
	}
	hasil := 1
	for i := 2; i <= n; i++ {
		hasil *= i
	}
	return hasil
}

// Fungsi untuk menghitung permutasi
func permutasi(n, r int) int {
	return faktorial(n) / faktorial(n-r)
}

// Fungsi untuk menghitung kombinasi
func kombinasi(n, r int) int {
	return faktorial(n) / (faktorial(r) * faktorial(n-r))
}

func main() {
	// Input
	var a, b, c, d int
	fmt.Println("Masukkan nilai a, b, c, dan d (pisahkan dengan spasi):")
	fmt.Scanf("%d %d %d %d", &a, &b, &c, &d)

	// Menghitung untuk a dan c
	P_a_c := permutasi(a, c)
	C_a_c := kombinasi(a, c)

	// Menghitung untuk b dan d
	P_b_d := permutasi(b, d)
	C_b_d := kombinasi(b, d)

	// Menampilkan hasil
	fmt.Printf("%d %d\n", P_a_c, C_a_c)
	fmt.Printf("%d %d\n", P_b_d, C_b_d)
}
