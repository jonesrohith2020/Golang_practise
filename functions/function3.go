package main

import "fmt"

func fibonacci(num int) {
	var a, b int
	for a, b = 0, 1; a <= num; a, b = b, a+b {
		fmt.Printf("%d\t", a)
	}
	fmt.Println()
}

func main() {
	var num int
	fmt.Print("Enter the number: ")
	fmt.Scanf("%d", &num)
	fibonacci(num)
}

// fibonacci series
