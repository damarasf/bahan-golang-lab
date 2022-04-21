package main

import "fmt"

func naikGaji(gajiAwal *int, gajiAkhir *int){
	*gajiAwal = *gajiAkhir
}

func main () {
	var gajiSekarang, expektasiGaji int

	fmt.Print("Masukkan gaji anda: ")
	fmt.Scan(&gajiSekarang)

	fmt.Print("Masukkan gaji yang diinginkan: ")
	fmt.Scan(&expektasiGaji)

	naikGaji(&gajiSekarang, &expektasiGaji)
	fmt.Println("Gaji anda sekarang: ", gajiSekarang)
}