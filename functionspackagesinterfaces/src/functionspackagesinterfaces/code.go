package main

import (
	"fmt"
	"libraries"
)

func main() {
	n1, l1 := libraries.FullNameNakedReturn("PF", "Vilquin")
	fmt.Printf("Fullname: %v, number of chars %v", n1, l1)
}
