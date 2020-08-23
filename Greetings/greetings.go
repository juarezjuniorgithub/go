package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println("Hi there! Type your name and age:")

	var myName string
	var age int = 45
	var golangOrg string = "Visit golang.org"

	fmt.Scanln(&myName)
	myName = strings.TrimSpace(myName)
	fmt.Scanln(&age)

	fmt.Printf("Hi %s, I'm Go!", myName)
	fmt.Println("")
	fmt.Printf("Your age is %d", age)
	fmt.Println("")
	fmt.Println(golangOrg)

}
