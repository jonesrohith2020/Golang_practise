package main

import "fmt"

func main() {
	apps := []string{"Facebook", "Instagram", "WhatsApp"}

	for i := 0; i < len(apps); i++ {
		fmt.Println(i, apps[i])
	}
}
