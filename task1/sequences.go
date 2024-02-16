package main

import (
	"fmt"
)

func main() {
	var menu []string

	menu = append(menu, "hamburger")
	menu = append(menu, "salad")

	for _, food_name := range menu {
		fmt.Println("Food:", food_name)
	}

	array := [5]int{4, 3, 2, 1, 0}
	for i, value := range array {
		fmt.Printf("This is %d and its index in the array is %d\n", value, i)
	}
}
