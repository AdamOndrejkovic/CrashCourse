package main

import (
	"fmt"
	"strconv"
)

//Define person struct
type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

// Greeting method(value receiver)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

//hasBirthday method (pointer receiver)
func (p *Person) hasBirthday() {
	p.age++
}

func main() {
	//Init person using struct
	person := Person{firstName: "Jane", lastName: "Doe", city: "Berlin", gender: "female", age: 30}
	personTwo := Person{"Bob", "Red", "Prague", "male", 28}

	fmt.Println(person)
	fmt.Println(personTwo)

	//Single field
	fmt.Println(person.firstName)
	//Change
	person.age = 31
	fmt.Println(person)

	person.hasBirthday()
	fmt.Println(person.greet())
}
