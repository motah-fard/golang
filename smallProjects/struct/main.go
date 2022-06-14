package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	// contactInfo contacInfo
	contactInfo
}

type contactInfo struct {
	email   string
	zipcode int
}

// card struct field defenition
// suit which is a st4ring
// value which is a defenition
// struct is like an object in js or it is like dictionary in python

func main() {
	// first way
	// alex := person{"Alex", "Anderson"} this is correct but the next line is correct too
	// second way
	// we do not want to have orders for it
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// fmt.Println(alex)
	// third way
	var alex person
	alex.firstName = "Alex"
	alex.lastName = "Anderson"
	// fmt.Println(alex)
	// "%+v" print out all the different field name from alex
	// fmt.Printf("%+v", alex)
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipcode: 44000,
		},
	}
	// fmt.Println(jim)
	// fmt.Printf("%+v", jim)
	jim.print()
	jimPointer := &jim
	jimPointer.updateName("jimmy")
	jim.print()

}

// structs with reciever functions
func (p person) print() {
	fmt.Printf("%+v", p)
}

// pass by value pointers
// go is a pass by value language: pass a value into a function:
// go copies that value and then the copy is made and then the copy is made available to the code
// *person : this is a type discription : it means we are working with a pointer to a person
// *pointerToPerson this is an operator : it means we want to manipulate the value the pointer is refrencing
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
