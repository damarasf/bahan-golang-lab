package main

import "fmt"

// Dari act2-lecture yang diberikan, cobalah untuk mempraktikan kemampuan kalian untuk menyelesaikan kasus ini
// multiple return values

// fungsi calculate akan mengembalikan hasil melakukan perhitungan berikut
// penjumlahan, pengurangan,perkalian,dan pembagian
func main() {
	sumResult, subtractResult, multiplyResult, divideResult := calculate(4, 4)
	fmt.Println("hasil penjumlahan", sumResult)
	fmt.Println("hasil pengurangan", subtractResult)
	fmt.Println("hasil perkalian", multiplyResult)
	fmt.Println("hasil pembagian", divideResult)
	fmt.Println(calculate(5, 5))
}

func calculate(Number1, Number2 int) (sumResult, subtractResult, multiplyResult, divideResult int) {
	// TODO: answer here
	
	return
}