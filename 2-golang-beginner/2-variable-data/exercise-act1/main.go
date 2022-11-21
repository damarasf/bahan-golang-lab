package main

import (
	"fmt"
)

// Dari lecture-act1 yang diberikan, cobalah untuk mempraktikan kemampuan kalian untuk menghitung volume tabung
// lengkapi dengan kode yang dibutuhkan untuk menghitung volume tabung

// Check Point:
// Menghitung volume tabung
// - Input: jari-jari, tinggi
// - Output: volume tabung

//rumus : V = π * r * r * t

// Contoh:
// Input:
// - Masukkan jari-jari alas tabung: 2
// - Masukkan tinggi tabung : 20
// Output:
// - Jadi volumenya adalah : 251.200012

func main() {
	// TODO: answer here
	var (
		r, t, V float64
	)
	const pi float64 = 3.14
	fmt.Print("Masukkan jari-jari : ")
	fmt.Scan(&r)
	fmt.Print("Masukkan tinggi : ")
	fmt.Scan(&t)

	V = pi * r * r * t
	fmt.Printf("Jadi volumenya adalah : %f", V)
}
