package main

import "fmt"

func main() {
	a, b := 2, 3

	// increment assignment
	a += b
	fmt.Println(a)

	// decrement assignment
	b -= 2
	fmt.Println(b)

	//multiplication assignment
	b *= 10
	fmt.Println(b)

	//division assignment
	b /= 5
	fmt.Println(b)

	// modules assignment
	a %= 3
	fmt.Println(a)
}
