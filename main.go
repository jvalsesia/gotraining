package main

import (
	"fmt"

	"github.com/jvalsesia/gotraining/pointers"
	"github.com/jvalsesia/gotraining/structtypes"
	"github.com/jvalsesia/gotraining/variables"
)

func main() {
	var option int

	fmt.Println("1 - variables: ")
	fmt.Println("2 - struct types: ")
	fmt.Println("3 - pointers: ")
	fmt.Println("4 - pointers 2: ")
	fmt.Scan(&option)

	switch option {
	case 1:
		variables.Exercise()
	case 2:
		structtypes.Exercise()
	case 3:
		pointers.Exercise()
	case 4:
		pointers.Exercise2()
	}

}
