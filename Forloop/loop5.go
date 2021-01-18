package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		if i > 5 {
			break
		}
		fmt.Printf("%d ", i)
	}
	fmt.Printf("\nloop is terminated if i>5")
}

// with break statement in for loop
