package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	jim := person{firstName: "Jim",
		lastName: "Grey",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		}}
	jim.updateFirstName("JonBonJones")
	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p person) updateFirstName(newFirstName string) {
	p.firstName = newFirstName
}
