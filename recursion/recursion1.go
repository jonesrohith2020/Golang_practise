package main

import "fmt"

func recursivefunction() {
	fmt.Println("Hello")
	recursivefunction()
}

func main() {
	recursivefunction()
}

//Infinite function calls
