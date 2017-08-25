package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

func (p person) printName() string { //(p person) indication method belongs to which type
	return p.firstName + " " + p.lastName
}

//Inheritance @ Embedding
type student struct {
	person
	class string
}

func main() {
	// p1 := person{"Niezwan", "Abdul Wahid", 20}
	// fmt.Println(p1.firstName)
	// fmt.Println(p1.lastName)
	// fmt.Println(p1.age)
	// fmt.Println(p1.printName())

	//Inheritance @ Embedding
	student := student{
		person: person{
			firstName: "Wan",
			lastName:  "AW",
			age:       23,
		},
		class: "3 Murni",
	}

	fmt.Println(student.printName())

}
