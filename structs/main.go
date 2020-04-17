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
	// jim as type person with fmt.Printf
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	fmt.Printf("%+v\n", jim)

	// pim as type person with fmt.Println
	pim := person{
		firstName: "Pim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "pim@gmail.com",
			zipCode: 95000,
		},
	}

	fmt.Println(pim)

	// First declare var kim as type person
	// then set variables and print with fmt.Printf
	var kim person

	kim.firstName = "Kim"
	kim.lastName = "Party"
	kim.contactInfo.email = "kim@gmail.com"
	kim.contactInfo.zipCode = 96000

	fmt.Printf("%+v\n", kim)

	// Use structs with receiver function
	var vin person

	vin.firstName = "Vin"
	vin.lastName = "Party"
	vin.contactInfo.email = "vin@gmail.com"
	vin.contactInfo.zipCode = 96000

	vinPointer := &vin
	fmt.Println(vinPointer)
	vinPointer.updateName("Vinny")
	vin.print()
}

func (pointerToPerson *person) updateName(newFirstname string)  {
	(*pointerToPerson).firstName = newFirstname
}

func (pointerToPerson *person) print() {
	fmt.Printf("%+v\n", (*pointerToPerson))
}
