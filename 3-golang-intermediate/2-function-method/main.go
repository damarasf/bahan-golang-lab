package main

import "fmt"

func main() {
	fmt.Println(IsEvenNumber(9))
}

func IsEvenNumber(number int) bool {
	if number%2 == 0 {
		return true
	}
	return false
}
