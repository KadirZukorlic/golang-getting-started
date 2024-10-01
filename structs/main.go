package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}
type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	nahla := person{
		firstName: "Nahla",
		lastName:  "Zukorlić",
		contactInfo: contactInfo{
			email:   "nahlazukorlic@gmail.com",
			zipCode: 36320,
		},
	}
	nahlaPointer := &nahla
	nahlaPointer.changeName("Hamza")
	nahla.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointerToPerson *person) changeName(newName string) {
	(*pointerToPerson).firstName = newName
}
