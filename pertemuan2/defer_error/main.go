package main

import "fmt"

func logging() {
	fmt.Println("Finish running the program")
}

func runAplication() {
	fmt.Println("Run application")
}

func main() {
	defer logging()
	runAplication()
}
