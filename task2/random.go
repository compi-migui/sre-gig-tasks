package main

import (
	"fmt"
	"math/rand"
)


func is_even(a int) bool {
	return a % 2 == 0
}

func main() {
	var n int = rand.Intn(101) // 1-100, both included

	switch {
	case n > 50:
		if is_even(n) {
			fmt.Println("It's closer to 100, and it's even!")
		} else {
			fmt.Println("It's closer to 100")
		}
	case n < 50:
		fmt.Println("It's closer to 0")
	case n == 50:
		fmt.Println("It's 50!")
	}
	fmt.Println(n)
}
