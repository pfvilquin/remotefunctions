package main

import (
	"fmt"
)

func main() {
	n1, l1 := FullNameNakedReturn("PF", "Vilquin")
	fmt.Printf("Fullname: %v, number of chars %v", n1, l1)
}

func FullNameNakedReturn(first, last string) (fullName string, length int) {
	fullName = first + " " + last
	length = len(fullName)
	return
}
