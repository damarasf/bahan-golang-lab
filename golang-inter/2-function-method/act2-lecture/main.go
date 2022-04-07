package main

import "fmt"

func main() {
	// memberikan nilai variabel sumResult dan multiplyResult
	// dengan nilai yang diberikan 10 dan 20
	sumResult, multiplyResult := calculate(10, 20)
	fmt.Printf("%d dan %d\n", sumResult, multiplyResult)
	fmt.Println(calculate(5, 5))
}

func calculate(angka1, angka2 int) (int, int) {
	sumResult := angka1 + angka2
	multiplyResult := angka1 * angka2
	return sumResult, multiplyResult
}
