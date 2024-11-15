package main

import (
	"fmt"

	"github.com/headfirstgo/keyboard"
)

func main() {
	var gpa [3]float64

	//go get github.com/headfirstgo/keyboard
	for i := 0; i < len(gpa); i++ {
		fmt.Print("Input float number : ")
		gpa[i], _ = keyboard.GetFloat()
	}
	for i, v := range gpa {
		fmt.Printf("GPA %d: %f\n", i, v)
	}
}
