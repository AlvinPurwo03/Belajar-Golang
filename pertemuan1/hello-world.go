package main

import "fmt"

func main() {
	var name string = "Alvin"

	switch name {
	case "Gera":
		fmt.Println("Namanya benar Gera & Ghema")
	case "Alvin":
		fmt.Println("Namanya Benar Alvin")
	default:
		fmt.Println("Bukan Gera dan bukan Alvin")
	}
}
