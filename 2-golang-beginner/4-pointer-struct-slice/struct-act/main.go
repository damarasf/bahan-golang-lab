package main

import "fmt"

type mahasiswa struct {
	nama  string
	kelas string
}

func main() {

	var data = map[string]mahasiswa{
		"51418626": {
			"Damara",
			"4IA18",
		},
		"51418647": {
			"Dinda",
			"4IA18",
		},
	}

	var search string
	fmt.Print("Masukkan nim mahasiswa yang ingin dicari : ")
	fmt.Scanln(&search)

	var result, ok = data[search]

	if ok {
		fmt.Println("Nama : ", result.nama)
		fmt.Println("Kelas : ", result.kelas)
	} else {
		fmt.Println("Nim tidak ditemukan")
	}
}
