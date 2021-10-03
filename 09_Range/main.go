package main

import "fmt"

func main() {
	ids := []int{1, 23, 42, 45, 54, 6, 2}

	//Loop through ids
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	//Not using index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	//Add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum", sum)

	// Range with map
	mainCities := map[string]string{"Belgium": "Brussels", "Slovakia": "Bratislava", "Denmark": "Copenhagen"}

	for k, v := range mainCities {
		fmt.Printf("%s: %s\n", k, v)
	}

	for k := range mainCities {
		fmt.Println("Country: " + k)
	}
}
