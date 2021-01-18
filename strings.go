package main

import "fmt"

func main() {
	var name string = "hi hello"
	var name1 string = "hihellohowareyou?"
	pi := 3.1415
	fmt.Println(len(name))
	fmt.Println(len(name1))
	fmt.Println(name1 + name)
	fmt.Println(name1 + "im fine") //append
	fmt.Printf("%T \n", name)      //gives the type of variable
	fmt.Printf("%b \n", 5)         //gives the binary
	fmt.Printf("%x \n", 15)        //hex values
	fmt.Printf("%e \n", pi)        //scientific notation
}
