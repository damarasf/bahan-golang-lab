package main

import (
	"fmt"
)

// Menghitung luas lingkaran
// - Input: jari-jari : nilai jari-jari lingkaran
// - Output: Jadi luasnya adalah : nilai luar lingkaran

// Contoh
// - Input: Masukkan jari : 10
// - Output: Jadi luasnya adalah : 314.000000000000

func main() {
	var (
		r    float64
		area float64
	)
	const pi float64 = 3.14
	fmt.Printf("Masukkan jari-jari : ")
	fmt.Scan(&r)
	area = pi * r * r
	fmt.Printf("Jadi luasnya adalah : %f", area)
}
