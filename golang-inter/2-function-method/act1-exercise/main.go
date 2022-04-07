package main

import "fmt"

// Dari act1-lecture yang diberikan, cobalah untuk mempraktikan kemampuan kalian untuk menyelesaikan kasus ini
// Fungsi printWord akan melakukan print satu persatu nilai parameter yang diterimanya
// Contoh nilai parameter yang diterima
// ("halo", "andi", "dan", "bagus")

//Outputnya :
// halo
// andi
// dan
// bagus

func main() {
	printWord("halo")
	printWord("halo", "selamat siang")
	printWord("halo", "andi", "dan", "bagus")
	printWord("mencoba", "variadic", "param", "pada", "go")
}

// TODO: answer here

