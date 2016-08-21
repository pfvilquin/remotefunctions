package main

import (
	"fmt"
)

func main() {
	var colors = [3]string{"Red", "Green", "Blue"}
	fmt.Println(colors)

	var colors2 = []string{"Red", "Green", "Blue"}
	fmt.Println(colors2)

	colors2 = append(colors2, "Purple", "White", "marine", "Black", "brown")
	fmt.Println(colors2)
	fmt.Println(cap(colors2))
	numbers := make([]int, 5, 50)
	for i:=0, j:=10 , i<5 {
		number[i++] = j 
	}
	
	

}
