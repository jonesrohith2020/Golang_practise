package main

import "fmt"

type Address struct {
	Name    string
	city    string
	Pincode int
}

func main() {

	var a Address
	fmt.Println(a)

	a1 := Address{"rohi", "Dgl", 123456}

	fmt.Println("Address1: ", a1)

	a2 := Address{Name: "keer", city: "madurai",
		Pincode: 624208}

	fmt.Println("Address2: ", a2)

	a3 := Address{Name: "Theni"}
	fmt.Println("Address3: ", a3)
}

//defines the declaration and definiton of the struct
