package variables

import (
	"fmt"
)

func Exercise() {
	// Declare variables that are set to their zero value.
	var age int
	var name string
	var legal bool

	// Display the value of those variables.
	fmt.Println(age)
	fmt.Println(name)
	fmt.Println(legal)

	// Declare variables and initialize.
	// Using the short variable declaration operator.
	month := 10
	dayOfWeek := "Tuesday"
	happy := true

	// Display the value of those variables.
	fmt.Println(month)
	fmt.Println(dayOfWeek)
	fmt.Println(happy)

	// Perform a type conversion.
	pi := float32(3.14)

	// Display the value of that variable.
	fmt.Printf("%T [%v]\n", pi, pi)

}
