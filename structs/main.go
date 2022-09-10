package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact contactInfo
}

func main() {
	jordan := person{
		firstName: "Jordan",
		lastName: "Daniel",
		contact: contactInfo{
			email: "jordandj2001@gmail.com",
			zipCode: 74000,
		},
	}

	pointerToJordan := &jordan
	pointerToJordan.updateFirstName("Jordannn")
	jordan.print()
}

func (pointer *person) updateFirstName(newFirstName string) {
	(*pointer).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}