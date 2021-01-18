package main

import "fmt"

func main() {
outer:
	for i := 0; i < 3; i++ {
		for j := 1; j <= 4; j++ {
			fmt.Printf("i = %d , j = %d\n", i, j) //labels can be used to break the outer loop from inside the inner for loop.
			if i == j {
				break outer
			}
		}

	}
}
