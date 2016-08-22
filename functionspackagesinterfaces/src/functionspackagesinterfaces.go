package main

import (
	"fmt"
)

func main() {
	n1, l1 := FullNameNakedReturn("PF", "Vilquin")
	fmt.Printf("Fullname: %v, number of chars %v", n1, l1)
}

func FullNameNakedReturn(first, last string) (full string, length int) {
	fullName = full + " " + last
	lenght = len(fullName)
	return
}
