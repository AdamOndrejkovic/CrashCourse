package main

import "fmt"

func main() {
	//Main  Types
	//string
	//bool
	//int
	//int int8 int16 int32 int64
	//uint uint8 uint16 uint32 uint64
	//byte - alias for unit8
	//rune - alias for int32
	//float32 float64
	//complex64 complex128

	//Using var

	var name string = "Adam"
	var age int32 = 21
	var isCool = true

	//Shorthand
	role := "Developer"
	size := 1.3

	city, country := "Esbjerg", "Denmark"

	fmt.Println(name, age, isCool, role, size, city, country)
	fmt.Printf("%T\n",age)


}
