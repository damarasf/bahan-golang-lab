package main

import (
	"fmt"
	"strings"
)

// Lengkapilah fungsi main dibawah ini agar program berjalan dengan benar

// Output yang diharapkan:
// halo Bintang Anugrah
// nama panggilan : Anugrah

type student struct {
	name  string
	grade int
}

func (s student) sayHello() {
	fmt.Println("halo", s.name)
}

func (s student) getNameAt(i int) string {
	return strings.Split(s.name, " ")[i-1]
}

func main() {
	//TODO: Answer here
	var s1 = student{"Bintang Anugrah", 10}
	s1.sayHello()

	var s2 = s1.getNameAt(2)
	fmt.Println("nama panggilan :", s2)
}
