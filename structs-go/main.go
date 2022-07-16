package main

import "fmt"

type contactInfo struct {
	email string
	zipcode int
}

type person struct {
	firstName string
	lastName string
	contactInfo
}


func main () {

	jim := person {
		firstName: "Jim",
		lastName: "Garyson",
		contactInfo : contactInfo {
			email: "jim@gmail.com",
			zipcode: 995995,
		},
	}
	
	jim.updateName("Jimmy")
	jim.print()
}

func (p person) print () {
	fmt.Printf("%+v\n",p)
}

func (pointerToPerson *person) updateName (newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}