package main

import (
	"fmt"
	"math"
)

func main() {
	var radius int
	var area float32

	fmt.Print("Enter radius: ")
	fmt.Scanf("%d", &radius)

	area = math.Pi * (float32(radius) * float32(radius))
	fmt.Printf("Area of circle: %.2f\n", area)
}
