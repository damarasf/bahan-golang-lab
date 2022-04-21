package main

import "fmt"

// Check Point:
// Diamond Pattern
// - Input: Size: 5
// - Output:
//         *
//        ***
//       *****
//      *******
//     *********
//    ***********
//     *********
//      *******
//       *****
//        ***
//         *

func main() {
	var i, j, k, l, size int

	fmt.Printf("Enter Size : ")
	fmt.Scan(&size)

	// TODO: answer here
	for i = size; i > 0; i-- {
		for j = 1; j <= i; j++ {
			fmt.Printf(" ")
		}
		for k = size; k > i-1; k-- {
			fmt.Printf("*")
		}
		for l = size; l > i; l-- {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}

	for i = 1; i <= size; i++ {
		for j = 1; j <= i; j++ {
			fmt.Printf(" ")
		}
		for k = size; k > i-1; k-- {
			fmt.Printf("*")
		}
		for l = size; l > i; l-- {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}
}
