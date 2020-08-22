package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println("Hello World")

	var myName string

	fmt.Scanln(&myName)
	myName = strings.TrimSpace(myName)

	fmt.Printf("Hi %s, I'm Go!", myName)

}
