package main

import (
	"fmt"
)

func main() {
	fmt.Println("Conditional Statement")

	var age int

	fmt.Print("Enter Your Age: ")
	fmt.Scan(&age)

	if age < 30 {
		fmt.Println("Not Allowed")
	} else {
		fmt.Println("Allowed")
	}

	var name string

	fmt.Print("Enter a Name: ")
	fmt.Scan(&name)

	if name == "Maruti" {
		fmt.Println(name)
	}

	EmptyStmt := ""
	fmt.Println(EmptyStmt)
}
