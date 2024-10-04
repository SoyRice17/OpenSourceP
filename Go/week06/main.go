package main

import (
	"fmt"
)

func main() {
	var i int
	var f float64
	var b bool
	var s string

	fmt.Println(f, b, s, i)
	fmt.Printf("%f%t%s%d", f, b, s, i) // zero value
	f = 2.7
	i = 3
	fmt.Print("\n\n", f > float64(i), "\n") // comparsion (true/false)
}
