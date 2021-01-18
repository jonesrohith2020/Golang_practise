package main

import "fmt"

type Employee struct {
	firstName, lastName string
	age, salary         int
}

func main() {

	emp := &Employee{"keerthu", "velu", 23, 6000}
	fmt.Println("First Name:", (*emp).firstName)
	fmt.Println("Age:", (*emp).age)
}

//pointer to sruct
