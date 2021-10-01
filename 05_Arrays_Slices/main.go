package main

import "fmt"

func main() {
	//Arrays
	var fruitArr [2]string

	//Assign Value
	fruitArr[0] = "Apple"
	fruitArr[1] = "Banana"

	//Declare and assign
	vegetableArr := [2]string{"carrot", "broccoli"}
	vegetableSlice := []string{"tomato", "potato","onions"}

	fmt.Println(fruitArr)
	fmt.Println(vegetableArr)
	fmt.Println(vegetableSlice)
	fmt.Println(len(fruitArr))
	fmt.Println(vegetableSlice[1:2])
}
