package main

import "fmt"

func factorial(num int) int {
	var fact = 1
	for i := 1; i <= num; i++ {
		fact = fact * i
	}
	return fact
}

func main() {
	var num, factorial_val int
	factorial_val = 1
	fmt.Print("Enter the number: ")
	fmt.Scanf("%d", &num)

	if num < 0 {
		fmt.Println("\n negative number")
	} else {
		factorial_val = factorial(num)
		fmt.Printf("Factorial of %d: %d\n", num, factorial_val)
	}
}

// finding factorial numbers
