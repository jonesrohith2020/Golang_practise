package main

import "fmt"

var num = 0

func count() {
	num++
	if num <= 10 {
		fmt.Println(num)
		count()
	}
}

func main() {
	count()
}

// recursive with finite number of times
