package main

import "fmt"

// int, float, string, bool, structs are value types (pointers are important)
// slices, maps, channels, pointers, and functions are reference types (pointers are not at as important?)

// Structs are similar to objects/classes. But instead of having values they just define types and variable names
type person struct {
	firstName string
	lastName  string
	contact   contactinfo
}

type contactinfo struct {
	email   string
	zipCode int
}

func (p *person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updateName(newFirstName string) {
	// go is a pass by value language, meaning whenever we pass a value into a function. Go will take that struct and create a copy of it, not use the reference.
	// &variable == give me the memory address of the value this variable is pointing at
	// *variable == give me the value this memory address is pointing at
	// func (pointerToThing *thing) funcname... { -> this is a type description, it means we're working with a pointer to a person
	//   *pointerToThing -> this is an operator, it means we will access or manipulate the value the pointer is referencing
	// }
	// so if we see a *variable normally, it means we want the value from a memory address. when it's in a function signature it means the function can only be called with a pointer, it's not an actual operator.

	// turn address to value with *variable
	// turn value to address with &variable
	p.firstName = newFirstName
}

func main() {
	calvin := person{
		firstName: "Calvin",
		lastName:  "Rice",
		contact: contactinfo{
			email:   "calvin@gmail.com",
			zipCode: 249012,
		},
	}

	calvin.updateName("Frankie")
	calvin.print()
	// if we don't define the values of a struct. it will default to a "zero value". Either literally 0 or an empty string or False.
	var toby person

	toby.print()
}
