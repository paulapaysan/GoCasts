package main

import "fmt"

type contactInfo struct {
	email string
	zip   int
}

type person struct {
	firstName   string
	lastName    string
	contactInfo // or contact contactInfo
}

func main() {
	// pam0 := person{"Pam", "Beesly"}
	// pam1 := person{firstName: "Pam", lastName: "Beesly"}
	var pam2 person
	pam2.firstName = "Pam"
	pam2.lastName = "Beesly"
	fmt.Printf("%+v", pam2)
	fmt.Println()

	jim := person{
		firstName: "Jim",
		lastName:  "Halpert",
		contactInfo: contactInfo{ // or contact : contactInfo {}
			email: "jim@gmail.com",
			zip:   18504,
		},
	}
	jimPointer := &jim
	jimPointer.updateName("Jimothy")
	jim.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
