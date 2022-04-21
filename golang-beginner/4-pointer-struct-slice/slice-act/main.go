package main

import "fmt"

func main() {
	var kursus = []string{"Go", "Python", "Ruby", "Java", "C++", "C#"}
	var kursus_saya = kursus[1:5]
	kursus_saya = append(kursus_saya, "Manajemen Informatika")

	fmt.Println("isi dari kursus adalah : ", kursus)
	fmt.Printf("panjang dari kursus adalah %d dan kapasitas %d\n", len(kursus), cap(kursus))

	fmt.Println("isi dari kursus_saya adalah : ", kursus_saya)
	fmt.Printf("panjang dari kursus_saya adalah %d dan kapasitas %d\n", len(kursus_saya), cap(kursus_saya))
}
