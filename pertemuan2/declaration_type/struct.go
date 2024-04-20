package main

import "fmt"

type User struct {
	ID    int
	Name  string
	Email string
}

func (u User) sayHello() {
	fmt.Println("Hello " + u.Name)
}

func main() {
	var john User

	john.ID = 1
	john.Name = "John"
	john.Email = "john@plabs.id"

	john.sayHello()
	//fmt.Printf("%#v\n", john)
}
