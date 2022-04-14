package main

import (
	"fmt"
)

// Dari act2-lecture yang diberikan, cobalah uji kemampuan kalian untuk mengimplementasikan program ini.
// Buatlah program kondisi mengenai penentuan total nilai UTS dengan bobot 30% dan UAS dengan bobot 70% termasuk Grade tertentu

// - Input: Nilai UTS dan UAS
//  Jika Total nilai UTS dan UAS <= 20, maka Grade E
//  Jika Total nilai UTS dan UAS > 20 dan <= 40, maka Grade D
//  Jika Total nilai UTS dan UAS > 40 dan <= 60, maka Grade C
//  Jika Total nilai UTS dan UAS > 60 dan <= 80, maka Grade B
//  Jika Total nilai UTS dan UAS > 80 dan <= 100, maka Grade A

// Output : hasil inputan nilai UTS, UAS, Total Nilai dan Grade

// Contoh:
// Input:
// - Masukkan nilai UTS : 90
// - Masukkan nilai UAS : 90
// Output:
// - Nilai UTS Anda : 90
// - Nilai UAS Anda : 90
// - Total Nilai : 90
// - Grade : A

func main () {
	var uts, uas, total int
	fmt.Printf("Masukkan nilai UTS : ")
	fmt.Scan(&uts)
	fmt.Printf("Masukkan nilai UAS : ")
	fmt.Scan(&uas)

	// TODO: answer here
	
}