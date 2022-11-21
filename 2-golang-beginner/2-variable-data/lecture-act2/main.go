package main

import (
	"fmt"
)

// Membuat kalkulator perhitungan sederhana
// - Input: angka pertama, angka kedua, operator
// - Output: hasil perhitungan

// Contoh
// - Input: Masukkan angka pertama : 10
// - Input: Masukkan angka kedua : 5
// - Input: Masukkan operator : +
// - Output: Hasilnya adalah : 15

// Masih banyak lagi operatornya, bisa diliat operatornya di sini:
// https://golang.org/ref/spec#Operators

func main() {
	var (
		a, b, c, d, e, f, g, h int
		operator               string
	)
	fmt.Print("Masukkan angka pertama : ")
	fmt.Scan(&a)
	fmt.Print("Masukkan angka kedua : ")
	fmt.Scan(&b)
	fmt.Print("Masukkan operator : ")
	fmt.Scan(&operator)

	if operator == "+" {
		c = a + b
		fmt.Println("Hasilnya adalah : ", c)
	} else if operator == "-" {
		d = a - b
		fmt.Println("Hasilnya adalah : ", d)
	} else if operator == "*" {
		e = a * b
		fmt.Println("Hasilnya adalah : ", e)
	} else if operator == "/" {
		f = a / b
		fmt.Println("Hasilnya adalah : ", f)
	} else if operator == "^" {
		g = a ^ b
		fmt.Println("Hasilnya adalah : ", g)
	} else if operator == "%" {
		h = a % b
		fmt.Println("Hasilnya adalah : ", h)
	} else {
		fmt.Println("Operator tidak dikenali")
	}
}
