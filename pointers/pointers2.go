// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare a struct type and create a value of this type. Declare a function
// that can change the value of some field in this struct type. Display the
// value before and after the call to your function.
package pointers

import "fmt"

// Add imports.

// Declare a type named user.
type user struct {
	name  string
	email string
	age   int
}

// Create a function that changes the value of one of the user fields.
/* add pointer parameter, add value parameter */
func changeAge(u *user, age int) {

	// Use the pointer to change the value that the
	// pointer points to.
	u.age = age
}

func Exercise2() {

	// Create a variable of type user and initialize each field.
	u1 := user{
		name:  "Benicio",
		email: "bvalsesia@gmail.com",
		age:   8,
	}

	// Display the value of the variable.
	fmt.Printf("%+v\n", u1)

	// Share the variable with the function you declared above.
	changeAge(&u1, 9)

	// Display the value of the variable.
	fmt.Printf("%+v\n", u1)

}
