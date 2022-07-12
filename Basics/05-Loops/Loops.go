package main

import (
	"fmt"
)


func main() {
	fmt.Println("Loops in Go")

	// Define a Slice
	Names := []string{"maruti", "Atul", "Mahesh", "Ravi"}
	fmt.Println(len(Names))

	// Basic For Loop
	for i := 0; i < len(Names); i++ {
		fmt.Println(Names[i])
	}

	// Append elements to slice
	Names = append(Names, "Pankaj", "Yogesh")
	surname := []string{"Gaikwad", "Jogdand"}

	// ... will unpack element of slice and append in first slice

	Names = append(Names, surname...)

	// For with condition
	i := 0

	for i < len(Names) {
		fmt.Println(Names[i])
		i++
	}

	// Infinite for loop
	for {
		if i < 10 {
			fmt.Println(i)
		} else {
			break
		}
		i++
	}

	// for each kind of loop

	for key, value := range Names {
		fmt.Println(key, value)
	}

	// Map
	Employees := make(map[int]string, 0)
	Employees[0] = "emp1"
	Employees[1] = "emp2"
	Employees[2] = "emp3"

	// For with range iterator
	for key, value := range Employees {
		fmt.Println(key, value)
	}

	// Even we give wrong key which is not present in map it will not through an error
	// but second value will be set to false
	// if key is in map then second value is true
	if val, ok := Employees[10]; ok {
		fmt.Println(val)
	}

}
