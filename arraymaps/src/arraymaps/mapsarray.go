package main

import (
	"fmt"
	//"math/rand"
	//"time"
)

type Dog struct {
	Breed  string
	Weight int
}

func main() {

	var colors = [3]string{"Red", "Green", "Blue"}
	fmt.Println(colors)

	var colors2 = []string{"Red", "Green", "Blue"}
	fmt.Println(colors2)

	colors2 = append(colors2, "Purple", "White", "marine", "Black", "brown", "Gold")
	fmt.Println(colors2)
	fmt.Println(cap(colors2))
	numbers := make([]int, 5, 50)

	for j, i := 10, 0; i < len(numbers); i, j = i+1, j+1 {
		numbers[i] = j
		j++
	}
	fmt.Println(numbers)

	poodle := Dog{"Poodle", 34}
	fmt.Println(poodle)
	fmt.Printf("%+v\n", poodle)

	/*	rand.Seed(time.Now().Unix())
		//dow := rand.Intn(6) + 1
		//fmt.Println("Day", dow)
		result := ""

		switch dow := rand.Intn(6) + 1; dow {
		case 1:
			result = "It's a Sunday"
		case 7:
			result = "It's a Saturday"
		default:
			result = "It's a Weekday"
		}
		fmt.Println("Day", result) */
	result := ""
	x := -0
	switch {
	case x < 0, x == 0:
		result = "negative or null"
		//fallthrough
	//case x == 0:
	//	result = "equal to 0"
	//fallthrough
	default:
		result = "Greater than 0"
	}
	fmt.Println(result)
}
