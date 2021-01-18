package main

import "fmt"

func main() {
	var a int = 5
	var b float32 = 10.05
	var e string = "hii"
	d := 12

	const (
		c         = 10
		f         = "hello"
		g float64 = 12.05345
	)
	isgirl := true
	x, y := "hii", "hello"
	ChangeValue(&d)
	fmt.Println(a)
	fmt.Println(&a)          //pointer
	fmt.Println("c+d=", c+d) //arithmetic
	fmt.Println(b)
	fmt.Println(!isgirl) //negotion boolean
	fmt.Println(e)       //string
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(f)
	fmt.Printf("%.2f \n", g) //printf example
	fmt.Println(x, ",", y)
}
func ChangeValue(d *int) { //pointer for changing the variable values in scope
	*d = 25
}
