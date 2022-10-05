package structtypes

import "fmt"

type user struct {
	name  string
	email string
	age   int
}

func Exercise() {

	u1 := user{
		name:  "Benicio",
		email: "bvalsesia@gmail.com",
		age:   8,
	}

	fmt.Printf("%+v\n", u1)

	// Display the field values.
	fmt.Println("Name", u1.name)
	fmt.Println("Email", u1.email)
	fmt.Println("Age", u1.age)

	// Declare a variable using an anonymous struct.
	u2 := struct {
		name  string
		email string
		age   int
	}{
		name:  "Julio",
		email: "jvalsesia@gmail.com",
		age:   46,
	}

	fmt.Printf("%+v\n", u2)

	// Display the field values.
	fmt.Println("Name", u2.name)
	fmt.Println("Email", u2.email)
	fmt.Println("Age", u2.age)
}
