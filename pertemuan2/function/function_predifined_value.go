package main

import "fmt"

func getName() (firstName, lastName string) {
	return "Alvin", "Purwo"
}

func main() {
	x, y := getName()
	fmt.Println(x, y)
}
