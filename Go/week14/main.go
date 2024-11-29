package main

import "fmt"

func main() {
	ages := make(map[string]int)

	var name string
	var age int

	for {
		fmt.Print("Enter name: (exit to 'q')")
		fmt.Scanln(&name)
		if name == "q" {
			break
		}
		fmt.Print("Enter age: ")
		fmt.Scanln(&age)
		ages[name] = age
	}
	for name, age := range ages {
		fmt.Printf("%s\t%d\n", name, age)
	}
}
