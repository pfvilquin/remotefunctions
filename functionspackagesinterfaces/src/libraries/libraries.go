package libraries

import "fmt"

func FullNameNakedReturn(first, last string) (fullName string, length int) {
	fullName = first + " " + last
	length = len(fullName)
	fmt.Println("Completed concatenation of firt and last name in Naked Return function")
	return
}
